package main

import (
	"context"
	"fmt"
	"log"
	"mxshop_srvs/order_srv/proto"

	"google.golang.org/grpc"
)

var OrderClient proto.OrderClient
var Conn *grpc.ClientConn

func Init() {
	var err error
	//使用grpc.Dial()进行拨号， grpc.WithInsecure()使用不安全的方式连接
	Conn, err = grpc.Dial("10.2.94.231:8083", grpc.WithInsecure())
	if err != nil {
		log.Panicln("连接失败", err)
	}
	OrderClient = proto.NewOrderClient(Conn)
}

func TestCreateShoppCart(userId, goodsId int32) {
	Rsp, err := OrderClient.CreateCarItem(context.Background(), &proto.CartItemRequest{
		UserId:  userId,
		GoodsId: goodsId,
		Nums:    3,
		Checked: true,
	})
	if err != nil {
		log.Fatal("加入购物车失败：", err)
		return
	}
	fmt.Println("id", Rsp.Id)
}

func TestCartItemList(userId int32) {
	Rsp, err := OrderClient.CartItemList(context.Background(), &proto.UserInfo{
		Id: userId,
	})
	if err != nil {
		log.Fatal("查询购物车列表失败", err)
		return
	}
	fmt.Println("用户购物车：", Rsp.Total, Rsp.Data)

}

func TestUpdateCartItem() {
	_, err := OrderClient.UpdateCartItem(context.Background(), &proto.CartItemRequest{
		Id:      7,
		Checked: true,
		Nums:    5,
	})
	if err != nil {
		log.Fatal("更新购物车失败", err)
		return
	}
	fmt.Println("更新购物车成功")
}

func DeleteCatItem() {
	_, err := OrderClient.DeleteCartItem(context.Background(), &proto.CartItemRequest{
		Id: 100,
	})
	if err != nil {
		log.Fatal("删除购物车记录失败", err)
		return
	}
	fmt.Println("删除成功")

}

func TestCreateOrder() {
	Rsp, err := OrderClient.CreateOrder(context.Background(), &proto.OrderRequest{
		UserId:  13,
		Address: "江苏省无锡市锡山区",
		Name:    "ice_moss",
		Mobile:  "17585610985",
		Post:    "麻烦店主给我发邮政快递",
	})
	if err != nil {
		log.Fatal("创建订单失败", err)
		return
	}
	fmt.Println("总价：", Rsp.Total)
}

func TestOrderList() {
	Rsp, err := OrderClient.OrderList(context.Background(), &proto.OrderFilterRequest{
		UserId:      13,
		Pages:       1,
		PagePerNums: 10,
	})
	if err != nil {
		log.Fatal("获取订单失败", err)
		return
	}
	for _, goods := range Rsp.Data {
		fmt.Println(goods.OrderSn)
	}
}

func TestOrderDetail() {
	Rsp, err := OrderClient.OrderDetail(context.Background(), &proto.OrderRequest{
		UserId: 6,
	})
	if err != nil {
		log.Fatal("'获取订单失败", err)
		return
	}
	fmt.Println(Rsp.OrderInfo, Rsp.Goods)
}

func TestUpdateStatus() {
	_, err := OrderClient.UpdateOrderStatus(context.Background(), &proto.OrderStatus{
		OrderSn: "202292085844639",
		Status:  "TRADE_FINISHED",
	})
	if err != nil {
		log.Fatal("更新订单失败", err)
		return
	}
	fmt.Println("更新订单状态成功")
}

func main() {
	Init()
	//TestCreateShoppCart(13, 425)
	//TestUpdateCartItem()
	//DeleteCatItem()
	//TestCreateOrder()
	//TestOrderList()
	TestOrderDetail()
	//TestUpdateStatus()
}
