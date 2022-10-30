package goods

import (
	"context"
	"fmt"

	"mxshop-api/goods-web/forms"
	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//HandleValidatorErr 表单验证错误处理返回
func HandleValidatorErr(c *gin.Context, err error) {
	fmt.Println(err.Error())
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": errs.Translate(global.Trans),
	})
}

//HandleGrpcErrToHttp grpc状态码转http
func HandleGrpcErrToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误",
				})
			}
		}
	}
}

func List(c *gin.Context) {
	var GoodsFilter proto.GoodsFilterRequest

	priceMin := c.DefaultQuery("pmin", "0")
	priceMinInt, _ := strconv.Atoi(priceMin)
	GoodsFilter.PriceMin = int32(priceMinInt)

	priceMax := c.DefaultQuery("pmax", "0")
	priceMaxInt, _ := strconv.Atoi(priceMax)
	GoodsFilter.PriceMax = int32(priceMaxInt)

	isHot := c.DefaultQuery("ih", "0")
	if isHot == "1" {
		GoodsFilter.IsHot = true
	}

	isNew := c.DefaultQuery("in", "0")
	if isNew == "1" {
		GoodsFilter.IsNew = true
	}

	isTab := c.DefaultQuery("it", "0")
	if isTab == "1" {
		GoodsFilter.IsTab = true
	}

	categoryId := c.DefaultQuery("c", "")
	categoryIdInt, _ := strconv.Atoi(categoryId)
	GoodsFilter.TopCategory = int32(categoryIdInt)

	Pages := c.DefaultQuery("pn", "0")
	PageInt, _ := strconv.Atoi(Pages)
	GoodsFilter.Pages = int32(PageInt)

	PageNums := c.DefaultQuery("pnum", "0")
	PageNumsInt, _ := strconv.Atoi(PageNums)
	GoodsFilter.PagePerNums = int32(PageNumsInt)

	keyWord := c.DefaultQuery("q", "")
	GoodsFilter.KeyWords = keyWord

	brandId := c.DefaultQuery("b", "0")
	brandIdInt, _ := strconv.Atoi(brandId)
	GoodsFilter.Brand = int32(brandIdInt)

	Rsp, err := global.GoodsSrvClient.GoodsList(context.Background(), &GoodsFilter)
	if err != nil {
		zap.S().Errorw("[list] 查找 【商品】失败", err.Error())
		HandleValidatorErr(c, err)
		return
	}

	var goodsList = make([]interface{}, 0)
	for _, value := range Rsp.Data {
		goodsList = append(goodsList, map[string]interface{}{
			"id":          value.Id,
			"name":        value.Name,
			"goods_brief": value.GoodsBrief,
			"desc":        value.GoodsDesc,
			"ship_free":   value.ShipFree,
			"images":      value.Images,
			"desc_images": value.DescImages,
			"front_image": value.GoodsFrontImage,
			"shop_price":  value.ShopPrice,
			"ctegory": map[string]interface{}{
				"id":   value.Category.Id,
				"name": value.Category.Name,
			},
			"brand": map[string]interface{}{
				"id":   value.Brand.Id,
				"name": value.Brand.Name,
				"logo": value.Brand.Logo,
			},
			"is_hot":  value.IsHot,
			"is_new":  value.IsNew,
			"on_sale": value.OnSale,
		})
	}

	var rspMap = make(map[string]interface{})
	rspMap["total"] = Rsp.Total
	rspMap["data"] = goodsList

	c.JSON(http.StatusOK, rspMap)
}

//New 添加商品
func New(c *gin.Context) {

	GoodsFrom := forms.GoodsFrom{}
	if err := c.ShouldBindJSON(&GoodsFrom); err != nil {
		HandleValidatorErr(c, err)
		return
	}

	Rsp, err := global.GoodsSrvClient.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
		Name:            GoodsFrom.Name,
		GoodsSn:         GoodsFrom.GoodsSn,
		Stocks:          GoodsFrom.Stocks,
		MarketPrice:     GoodsFrom.MarketPrice,
		ShopPrice:       GoodsFrom.ShopPrice,
		GoodsBrief:      GoodsFrom.GoodsBrief,
		ShipFree:        *GoodsFrom.ShipFree,
		Images:          GoodsFrom.Images,
		DescImages:      GoodsFrom.DescImages,
		GoodsFrontImage: GoodsFrom.FrontImage,
		CategoryId:      GoodsFrom.CategoryId,
		BrandId:         GoodsFrom.Brand,
	})
	if err != nil {
		HandleGrpcErrToHttp(err, c)
		return
	}

	//如何设置库存
	//TODO 商品的库存
	c.JSON(http.StatusOK, Rsp)
}

func Detail(c *gin.Context) {
	goodsId := c.Param("id")
	goodsIdInt, err := strconv.ParseInt(goodsId, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	Rsp, err := global.GoodsSrvClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: int32(goodsIdInt),
	})
	if err != nil {
		HandleGrpcErrToHttp(err, c)
		return
	}

	//TODO 去库存服务里查询库存

	RspGoods := map[string]interface{}{
		"id":          Rsp.Id,
		"name":        Rsp.Name,
		"goods_brief": Rsp.GoodsBrief,
		"desc":        Rsp.GoodsDesc,
		"ship_free":   Rsp.ShipFree,
		"images":      Rsp.Images,
		"desc_images": Rsp.DescImages,
		"front_image": Rsp.GoodsFrontImage,
		"shop_price":  Rsp.ShopPrice,
		"ctegory": map[string]interface{}{
			"id":   Rsp.Category.Id,
			"name": Rsp.Category.Name,
		},
		"brand": map[string]interface{}{
			"id":   Rsp.Brand.Id,
			"name": Rsp.Brand.Name,
			"logo": Rsp.Brand.Logo,
		},
		"is_hot":  Rsp.IsHot,
		"is_new":  Rsp.IsNew,
		"on_sale": Rsp.OnSale,
	}

	c.JSON(http.StatusOK, RspGoods)
}

//Delete 删除商品
func Delete(c *gin.Context) {
	goodsId := c.Param("id")
	goodsIdInt, err := strconv.ParseInt(goodsId, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	_, err = global.GoodsSrvClient.DeleteGoods(context.Background(), &proto.DeleteGoodsInfo{
		Id: int32(goodsIdInt),
	})
	if err != nil {
		HandleGrpcErrToHttp(err, c)
	}

	c.Status(http.StatusOK)
	return
}

//Stocks 获取库存
func Stocks(c *gin.Context) {
	goodsId := c.Param("id")
	_, err := strconv.ParseInt(goodsId, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	//TODO  商品库存服务
	return
}

//UpdateStatus 更新商品状态
func UpdateStatus(c *gin.Context) {
	//获取表单数据
	goodsStatusForm := forms.GoodsStatusForm{}
	if err := c.ShouldBind(&goodsStatusForm); err != nil {
		HandleValidatorErr(c, err)
		return
	}

	goodsId := c.Param("id")
	id, err := strconv.ParseInt(goodsId, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	//获取商品对应的品牌和分类
	rsp, err := global.GoodsSrvClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: int32(id),
	})
	if err != nil {
		HandleGrpcErrToHttp(err, c)
	}

	createGoods := proto.CreateGoodsInfo{
		Id:              int32(id),
		CategoryId:      rsp.CategoryId,
		BrandId:         rsp.Brand.Id,
		Name:            rsp.Name,
		GoodsSn:         rsp.GoodsSn,
		MarketPrice:     rsp.MarketPrice,
		ShopPrice:       rsp.ShopPrice,
		GoodsBrief:      rsp.GoodsBrief,
		ShipFree:        rsp.ShipFree,
		Images:          rsp.Images,
		DescImages:      rsp.DescImages,
		GoodsFrontImage: rsp.GoodsFrontImage,
		IsNew:           *goodsStatusForm.IsNew,
		IsHot:           *goodsStatusForm.IsHot,
		OnSale:          *goodsStatusForm.OnSale,
	}

	_, err = global.GoodsSrvClient.UpdateGoods(context.Background(), &createGoods)
	if err != nil {
		HandleGrpcErrToHttp(err, c)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
	})
}

//Update 更新商品信息
func Update(c *gin.Context) {
	GoodsFrom := forms.GoodsFrom{}
	if err := c.ShouldBind(&GoodsFrom); err != nil {
		HandleValidatorErr(c, err)
		return
	}

	goodsId := c.Param("id")
	goodsIdInt, err := strconv.ParseInt(goodsId, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	_, err = global.GoodsSrvClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:              int32(goodsIdInt),
		Name:            GoodsFrom.Name,
		GoodsSn:         GoodsFrom.GoodsSn,
		Stocks:          GoodsFrom.Stocks,
		MarketPrice:     GoodsFrom.MarketPrice,
		ShopPrice:       GoodsFrom.ShopPrice,
		GoodsBrief:      GoodsFrom.GoodsBrief,
		ShipFree:        *GoodsFrom.ShipFree,
		Images:          GoodsFrom.Images,
		DescImages:      GoodsFrom.DescImages,
		GoodsFrontImage: GoodsFrom.FrontImage,
		CategoryId:      GoodsFrom.CategoryId,
		BrandId:         GoodsFrom.Brand,
	})
	if err != nil {
		HandleGrpcErrToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
	})
}
