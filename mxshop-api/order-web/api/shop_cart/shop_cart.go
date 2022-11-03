package shop_cart

import (
	"context"
	"net/http"
	"strconv"

	"mxshop-api/order-web/api"
	"mxshop-api/order-web/forms"
	"mxshop-api/order-web/global"
	"mxshop-api/order-web/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//List 用户获取购物车列表
func List(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	ShopCartRsp, err := global.OrderSrvClient.CartItemList(context.Background(), &proto.UserInfo{
		Id: int32(userId.(uint)),
	})
	if err != nil {
		zap.S().Info("[List] 获取【购物车列表】失败", err)
		api.HandleGrpcErrToHttp(err, ctx)
		return
	}
	var goodsIdS []int32
	for _, value := range ShopCartRsp.Data {
		goodsIdS = append(goodsIdS, value.GoodsId)
	}

	if len(goodsIdS) == 0 {
		zap.S().Info("购物车数据为空")
		ctx.JSON(http.StatusOK, gin.H{
			"total": 0,
		})
		return
	}

	GoodsRsp, err := global.GoodsSrvClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: goodsIdS,
	})
	if err != nil {
		zap.S().Info("[List] 批量查询 【商品服务】失败")
		api.HandleGrpcErrToHttp(err, ctx)
		return
	}

	goodsList := make([]interface{}, 0)
	ReqMap := make(map[string]interface{})

	for _, item := range ShopCartRsp.Data {
		for _, goods := range GoodsRsp.Data {
			if item.GoodsId == goods.Id {
				tmpMap := map[string]interface{}{}
				tmpMap["id"] = item.Id
				tmpMap["goods_id"] = goods.Id
				tmpMap["goods_name"] = goods.Name
				tmpMap["goods_price"] = goods.ShopPrice
				tmpMap["goods_images"] = goods.Images
				tmpMap["nums"] = item.Nums
				tmpMap["checked"] = item.Checked

				goodsList = append(goodsList, tmpMap)
			}

		}
	}

	ReqMap["data"] = goodsList
	ctx.JSON(http.StatusOK, ReqMap)

}

//CreateCarItem 添加商品到购物车
func CreateCarItem(ctx *gin.Context) {
	var cartItem forms.ShopCartForm
	if err := ctx.ShouldBindJSON(&cartItem); err != nil {
		zap.S().Info("获取表单失败", err)
		api.HandleValidatorErr(ctx, err)
		return
	}

	//查询商品是否存在
	_, err := global.GoodsSrvClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: cartItem.GoodsId,
	})
	if err != nil {
		zap.S().Info("查询商品信息失败", err)
		api.HandleValidatorErr(ctx, err)
		return
	}

	//查询库存
	InvGoods, err := global.InventorySrvClient.InvDetail(context.Background(), &proto.GoodsInventoryInfo{
		GoodsId: cartItem.GoodsId,
	})
	if err != nil {
		zap.S().Info("查询库存失败", err)
		api.HandleGrpcErrToHttp(err, ctx)
		return
	}

	//判断库存是否充足
	if cartItem.Nums > InvGoods.Num {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "库存不足",
		})
		return
	}

	userId, _ := ctx.Get("userId")
	CartItemRsp, err := global.OrderSrvClient.CreateCarItem(context.Background(), &proto.CartItemRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: cartItem.GoodsId,
		Nums:    cartItem.Nums,
	})
	if err != nil {
		zap.S().Info("加入购物车失败", err)
		api.HandleGrpcErrToHttp(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": CartItemRsp.Id,
	})
}

func UpdateCarItem(ctx *gin.Context) {
	ShopCartUpdateForm := forms.UpdateShopCartForm{}
	if err := ctx.ShouldBindJSON(&ShopCartUpdateForm); err != nil {
		zap.S().Info("获取表单失败", err)
		api.HandleValidatorErr(ctx, err)
		return
	}

	CartItem := ctx.Param("id")
	CartItemId, err := strconv.Atoi(CartItem)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "url格式出错",
		})
		return
	}

	userId, _ := ctx.Get("userId")

	request := proto.CartItemRequest{
		GoodsId: int32(CartItemId),
		UserId:  int32(userId.(uint)),
		Nums:    ShopCartUpdateForm.Nums,
		Checked: false,
	}
	if ShopCartUpdateForm.Checked != nil {
		request.Checked = *ShopCartUpdateForm.Checked
	}

	_, err = global.OrderSrvClient.UpdateCartItem(context.Background(), &request)
	if err != nil {
		zap.S().Info("更新购物车记录失败", err)
		api.HandleGrpcErrToHttp(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
}

//DeleteCarItem 移除购物车
func DeleteCarItem(ctx *gin.Context) {
	CartItem := ctx.Param("id")
	CartItemId, err := strconv.Atoi(CartItem)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "url格式出错",
		})
		return
	}

	userId, _ := ctx.Get("userId")

	_, err = global.OrderSrvClient.DeleteCartItem(context.Background(), &proto.CartItemRequest{
		GoodsId: int32(CartItemId),
		UserId:  int32(userId.(uint)),
	})
	if err != nil {
		zap.S().Info("移除购物车记录失败", err)
		api.HandleGrpcErrToHttp(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "移除购物车成功",
	})
}
