package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"mxshop_srvs/inventory_srv/global"
	"mxshop_srvs/inventory_srv/model"
	"mxshop_srvs/inventory_srv/proto"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-redsync/redsync/v4"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

//InventoryServer 库存服务
type InventoryServer struct{}

//SetInv 设置库存或者库存更新
func (i *InventoryServer) SetInv(c context.Context, req *proto.GoodsInventoryInfo) (*empty.Empty, error) {
	var inventory model.Inventory

	global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inventory)
	inventory.Goods = req.GoodsId
	inventory.Stocks = req.Num
	global.DB.Save(&inventory)

	return &empty.Empty{}, nil
}

//InvDetail 查询库存
func (i *InventoryServer) InvDetail(c context.Context, req *proto.GoodsInventoryInfo) (*proto.GoodsInventoryInfo, error) {
	var inventory model.Inventory
	if result := global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inventory); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "库存信息不存在")
	}

	return &proto.GoodsInventoryInfo{
		GoodsId: inventory.Goods,
		Num:     inventory.Stocks,
	}, nil
}

//使用mutex只能在单个服务实例上进行并发锁
//当多个服务实例运行至不同服务器且同时访问一个数据库时
//mutex的锁就不起作用了，此时可以使用mysql的锁来完成分布式锁的功能
//var ms sync.Mutex

//Sell 扣减库存,涉及事务逻辑，执行的逻辑必须全部成功或者全部失败并且失败后数据可恢复,不能中途失败
func (i *InventoryServer) Sell(ctx context.Context, req *proto.SellInfo) (*empty.Empty, error) {

	socktDetail := model.StockSellDetail{
		OrderSn: req.OrderSn,
		Status:  1,
	}

	//并发情况下可能会出现超买，需要使用锁来将并发串行化
	tx := global.DB.Begin()
	var goodsdiscr []model.GoodsDetail
	var mutexs []*redsync.Mutex
	for _, goodsInfo := range req.GoodsInfo {
		goodsdiscr = append(goodsdiscr, model.GoodsDetail{
			Goods: goodsInfo.GoodsId,
			Num:   goodsInfo.Num,
		})

		var inventory model.Inventory
		mutex := global.Rs.NewMutex(fmt.Sprintf("goods_%d", goodsInfo.GoodsId))

		if err := mutex.Lock(); err != nil {
			return nil, status.Errorf(codes.Internal, "获取redis分布式锁异常")
		}

		if result := global.DB.Where(&model.Inventory{Goods: goodsInfo.GoodsId}).First(&inventory); result.RowsAffected == 0 {
			//失败进行事务回滚
			tx.Rollback()
			return nil, status.Errorf(codes.InvalidArgument, "库存信息不存在")
		}

		if inventory.Stocks < goodsInfo.Num {
			//失败进行事务回滚
			tx.Rollback()
			return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
		}
		inventory.Stocks -= goodsInfo.Num
		tx.Save(&inventory)

		mutexs = append(mutexs, mutex)

		//if ok, err := mutex.Unlock(); !ok || err != nil {
		//	return nil, status.Errorf(codes.Internal, "释放redis分布式锁异常")
		//}

	}
	socktDetail.Detail = goodsdiscr
	if result := tx.Create(&socktDetail); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "保存扣减库存历史数据失败")
	}

	//提交事务
	tx.Commit()

	for _, mutex := range mutexs {
		if ok, err := mutex.Unlock(); !ok || err != nil {
			return nil, status.Errorf(codes.Internal, "释放redis分布式锁异常")
		}
	}
	return &empty.Empty{}, nil
}

//Reback 同样涉及事务逻辑；归还库存：1.超时归还  2.订单创建失败，归还库存扣减  3.手动归还
func (i *InventoryServer) Reback(ctx context.Context, req *proto.SellInfo) (*empty.Empty, error) {
	//给服务上锁
	//ms.Lock()
	//事务开始
	tx := global.DB.Begin()
	for _, goodsInfo := range req.GoodsInfo {

		var inventory model.Inventory
		for {
			if result := global.DB.Where(&model.Inventory{Goods: goodsInfo.GoodsId}).First(&inventory); result.RowsAffected == 0 {
				//失败进行事务回滚
				tx.Rollback()
				return nil, status.Errorf(codes.InvalidArgument, "库存信息不存在")
			}
			inventory.Stocks += goodsInfo.Num
			if result := tx.Model(&model.Inventory{}).Select("Stocks", "Version").Where("goods = ? and version = ?",
				goodsInfo.GoodsId, inventory.Version).Updates(model.Inventory{Stocks: inventory.Stocks, Version: inventory.Version + 1}); result.RowsAffected == 0 {
				zap.S().Info("库存归还失败")
			} else {
				break
			}
		}
	}
	//提交事务
	tx.Commit()
	//释放锁
	//ms.Unlock()

	return &empty.Empty{}, nil
}

func AutoReback(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	type OrderInfo struct {
		OrderSn string
	}
	for i := range msgs {
		//生产者成功扣除库存后，将订单信息发送至qm,消费者inventory并做归还库存操作
		//归还商品，应该知道每一件商品归还多少件，重复归还商品问题需要解决
		//所以此接口需要幂等性，不能因为消息多次重复发送，导致商品重复归还
		//解决方案：新建一张表，记录详细的单订扣减记录，以及归还细节

		var orderInfo OrderInfo
		err := json.Unmarshal(msgs[i].Body, &orderInfo)
		if err != nil {
			zap.S().Errorf("解析json失败：%s", msgs[i].Body)
			return consumer.ConsumeSuccess, nil
		}

		//库存扣减记录
		var stockSellDetail model.StockSellDetail
		//开始事务
		tx := global.DB.Begin()

		//获取需要归还的库存
		if result := tx.Model(&model.StockSellDetail{}).Where(&model.StockSellDetail{OrderSn: orderInfo.OrderSn, Status: 1}).First(&stockSellDetail); result.RowsAffected == 0 {
			return consumer.ConsumeSuccess, nil
		}
		//如果查询到，各个归还
		for _, orderGoods := range stockSellDetail.Detail {
			//先查询inventory, update语句update xxx set stocks=stocks+2 当多个并发进入mysql会自动锁住，更安全
			if result := tx.Model(&model.Inventory{}).Where(&model.Inventory{Goods: orderGoods.Goods}).Update("stocks", gorm.Expr("stocks+?", orderGoods.Num)); result.RowsAffected == 0 {
				tx.Rollback()
				//过一段时间重新消费ConsumeRetryLater
				return consumer.ConsumeRetryLater, nil
			}

			//更新状态
			if result := tx.Model(&model.StockSellDetail{}).Where(&model.StockSellDetail{OrderSn: stockSellDetail.OrderSn}).Update("status", 2); result.RowsAffected == 0 {
				tx.Rollback()
				return consumer.ConsumeRetryLater, nil
			}
		}
		tx.Commit()

		return consumer.ConsumeSuccess, nil
	}
	return consumer.ConsumeSuccess, nil
}
