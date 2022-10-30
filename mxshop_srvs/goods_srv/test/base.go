package main

import (
	"log"
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/grpc"
)

var BrandClient proto.GoodsClient
var Conn *grpc.ClientConn

func Init() {
	var err error
	//使用grpc.Dial()进行拨号， grpc.WithInsecure()使用不安全的方式连接
	Conn, err = grpc.Dial("10.2.121.49:8081", grpc.WithInsecure())
	if err != nil {
		log.Panicln("连接失败", err)
	}
	BrandClient = proto.NewGoodsClient(Conn)
}

func main() {
	Init()
	//TestGetCategoryAllList()
	//TestGetSubCategory()
	//TestCreateCategory()
	//TestDeleteCategory()

	//TestUpdateCategory()

	//TestCategoryBrandList()
	//TestGetCategoryBrandList()
	//TestCreateBrand()

	//TestCreateCategoryBrand()
	//TestDeleteCategoryBrand()
	//TestUpdateCategoryBrand()
	//TestGetCategoryBrandList()

	//TestGoodsList()

	//TestGetGoodsDetail()

	//TestGetGoodsDetail()

	//TestBatchGetGoods()

	//TestCreateGoods()
	//TestUpdateGoods()

	//TestUpdateBanner()

	TestGetCategoryBrandList()

	Conn.Close()

}
