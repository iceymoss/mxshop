package order

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"mxshop-api/order-web/api"
	"mxshop-api/order-web/forms"
	"mxshop-api/order-web/global"
	"mxshop-api/order-web/models"
	"mxshop-api/order-web/proto"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//List 订单列表
func List(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	claimsInfo, _ := ctx.Get("claims")
	claims := claimsInfo.(*models.CustomClaims)

	Request := proto.OrderFilterRequest{}

	Pages := ctx.DefaultQuery("pn", "0")
	PageInt, _ := strconv.Atoi(Pages)
	Request.Pages = int32(PageInt)

	PageNums := ctx.DefaultQuery("pnum", "0")
	PageNumsInt, _ := strconv.Atoi(PageNums)
	Request.PagePerNums = int32(PageNumsInt)

	if claims.AuthorityId == 1 {
		Request.UserId = int32(userId.(uint))
	}

	//对商品列表进行限流
	e, b := sentinel.Entry("order_list", sentinel.WithTrafficType(base.Inbound))
	if b != nil {
		fmt.Println("限流了")
		ctx.JSON(http.StatusTooManyRequests, gin.H{
			"msg": "请求过于繁忙，请稍后重试",
		})
		return
	}

	Rsp, err := global.OrderSrvClient.OrderList(context.WithValue(context.Background(), "ginContext", ctx), &Request)
	if err != nil {
		zap.S().Info("[List] 获取【订单列表】失败")
		api.HandleGrpcErrToHttp(err, ctx)
		return
	}

	//退出限流
	e.Exit()

	OrderList := make([]interface{}, 0)
	for _, Item := range Rsp.Data {
		ItemMap := map[string]interface{}{}
		ItemMap["id"] = Item.Id
		ItemMap["name"] = Item.Name
		ItemMap["total"] = Item.Total
		ItemMap["userId"] = Item.UserId
		ItemMap["status"] = Item.Status
		ItemMap["orderSn"] = Item.OrderSn
		ItemMap["address"] = Item.Address
		ItemMap["mobile"] = Item.Mobile
		ItemMap["post"] = Item.Post
		ItemMap["pay_type"] = Item.PayType
		ItemMap["add_time"] = Item.AddTime
		OrderList = append(OrderList, Item)
	}

	ReMap := gin.H{
		"total": Rsp.Total,
		"data":  OrderList,
	}
	ctx.JSON(http.StatusOK, ReMap)
}

//CreaOrder 新建订单
func CreatOrder(ctx *gin.Context) {
	var OrderForm forms.OrderForms
	if err := ctx.ShouldBindJSON(&OrderForm); err != nil {
		zap.S().Info("获取表单失败")
		api.HandleValidatorErr(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")

	//对商品列表进行限流
	e, b := sentinel.Entry("create_order", sentinel.WithTrafficType(base.Inbound))
	if b != nil {
		fmt.Println("限流了")
		ctx.JSON(http.StatusTooManyRequests, gin.H{
			"msg": "请求过于繁忙，请稍后重试",
		})
		return
	}

	Rsp, err := global.OrderSrvClient.CreateOrder(context.WithValue(context.Background(), "ginContext", ctx), &proto.OrderRequest{
		UserId:  int32(userId.(uint)),
		Name:    OrderForm.Name,
		Address: OrderForm.Address,
		Mobile:  OrderForm.Mobile,
		Post:    OrderForm.Post,
	})
	if err != nil {
		zap.S().Info("新建订单失败")
		api.HandleGrpcErrToHttp(err, ctx)
		return
	}

	//退出限流
	e.Exit()

	//TODO 此时的逻辑跳转至支付宝支付页面，可通过web层或是srv层返回支付宝支付URL
	ctx.JSON(http.StatusOK, gin.H{
		"id": Rsp.Id,
	})
}

//DetailOrder 获取订单详情
func DetailOrder(ctx *gin.Context) {
	orderId := ctx.Param("id")
	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		zap.S().Info("url格式不正确")
		ctx.Status(http.StatusBadRequest)
		return
	}
	userId, _ := ctx.Get("userId")
	claimsInfo, _ := ctx.Get("claims")
	claims := claimsInfo.(*models.CustomClaims)

	OrderRequest := proto.OrderRequest{}
	OrderRequest.Id = int32(orderIdInt)

	if claims.AuthorityId == 1 {
		OrderRequest.UserId = int32(userId.(uint))
	}

	Rsp, err := global.OrderSrvClient.OrderDetail(context.WithValue(context.Background(), "ginContext", ctx), &OrderRequest)
	if err != nil {
		zap.S().Info("获取订单详情失败")
		api.HandleGrpcErrToHttp(err, ctx)
		return
	}

	reMap := make(map[string]interface{})
	reMap["id"] = Rsp.OrderInfo.Id
	reMap["orderSn"] = Rsp.OrderInfo.OrderSn
	reMap["name"] = Rsp.OrderInfo.Name
	reMap["userId"] = Rsp.OrderInfo.UserId
	reMap["status"] = Rsp.OrderInfo.Status
	reMap["payType"] = Rsp.OrderInfo.PayType
	reMap["post"] = Rsp.OrderInfo.Post
	reMap["address"] = Rsp.OrderInfo.Address
	reMap["mobile"] = Rsp.OrderInfo.Mobile
	reMap["total"] = Rsp.OrderInfo.Total
	reMap["addTime"] = Rsp.OrderInfo.AddTime

	GoodsList := make([]interface{}, 0)
	for _, goods := range Rsp.Goods {
		goodsItem := map[string]interface{}{}
		goodsItem["id"] = goods.Id
		goodsItem["name"] = goods.GoodsName
		goodsItem["goodsId"] = goods.GoodsId
		goodsItem["image"] = goods.GoodsImage
		goodsItem["nums"] = goods.Nums
		goodsItem["price"] = goods.GoodsPrice
		goodsItem["orderId"] = goods.OrderId
		GoodsList = append(GoodsList, goodsItem)
	}

	reMap["goods"] = GoodsList
	ctx.JSON(http.StatusOK, reMap)
}

func UpdateOrder(c *gin.Context) {

}
