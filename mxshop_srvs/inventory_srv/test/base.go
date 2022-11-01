package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"mxshop_srvs/inventory_srv/proto"

	"google.golang.org/grpc"
)

var InventoryClient proto.InventoryClient
var Conn *grpc.ClientConn

func Init() {
	var err error
	//使用grpc.Dial()进行拨号， grpc.WithInsecure()使用不安全的方式连接
	Conn, err = grpc.Dial("10.2.94.231:8082", grpc.WithInsecure())
	if err != nil {
		log.Panicln("连接失败", err)
	}
	InventoryClient = proto.NewInventoryClient(Conn)
}

func TestSetInv(goodsId, num int32) {
	_, err := InventoryClient.SetInv(context.Background(), &proto.GoodsInventoryInfo{
		GoodsId: goodsId,
		Num:     num,
	})
	if err != nil {
		log.Fatal("设置库存失败", err)
	}
	fmt.Println("设置库存成功")
}

func TestInDetail(goodsId int32) {
	Rsp, err := InventoryClient.InvDetail(context.Background(), &proto.GoodsInventoryInfo{
		GoodsId: goodsId,
	})
	if err != nil {
		log.Fatal("查询库存失败", err)
	}
	fmt.Println("库存信息", Rsp)
}

func TestSell(wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := InventoryClient.Sell(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInventoryInfo{
			{GoodsId: 421, Num: 6},
		},
	})
	if err != nil {
		log.Fatal("扣减库存失败", err)
	}
	fmt.Println("扣减库存成功")
}

func TestReback(wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := InventoryClient.Reback(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInventoryInfo{
			{GoodsId: 421, Num: 1},
		},
	})
	if err != nil {
		log.Fatal("归还库存失败", err)
	}
	fmt.Println("归还库存成功")
}

func main() {
	Init()
	//TestSetInv(421, 100)
	//TestInDetail(421)
	//TestSell()
	//TestReback()
	var wg sync.WaitGroup
	wg.Add(35)
	for i := 0; i < 35; i++ {
		//go TestSell(&wg)
		go TestReback(&wg)
	}
	wg.Wait()

	Conn.Close()

}
