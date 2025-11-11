package banner

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"mxshop-api/goods-web/forms"
	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/proto"
	"mxshop-api/user-web/api"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
)

// BannerList 获取轮播图列表
func BannerList(c *gin.Context) {
	Rsp, err := global.GoodsSrvClient.BannerList(context.WithValue(context.Background(), "ginContext", c), &empty.Empty{})
	if err != nil {
		api.HandleGrpcErrToHttp(err, c)
		return
	}

	data := make([]interface{}, 0)
	for _, value := range Rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["index"] = value.Index
		reMap["image"] = value.Image
		reMap["url"] = value.Url

		data = append(data, reMap)
	}
	c.JSON(http.StatusOK, data)
}

// NewBanner 添加轮播图
func NewBanner(c *gin.Context) {
	BannerForm := forms.BannerForm{}
	if err := c.ShouldBindJSON(&BannerForm); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}

	Rsp, err := global.GoodsSrvClient.CreateBanner(context.WithValue(context.Background(), "ginContext", c), &proto.BannerRequest{
		Index: int32(BannerForm.Index),
		Image: BannerForm.Image,
		Url:   BannerForm.Url,
	})
	if err != nil {
		api.HandleGrpcErrToHttp(err, c)
		return
	}
	response := make(map[string]interface{})
	response["id"] = Rsp.Id
	response["index"] = Rsp.Index
	response["url"] = Rsp.Url
	response["image"] = Rsp.Image

	c.JSON(http.StatusOK, response)
}

// DeleteBanner 删除轮播图
func DeleteBanner(c *gin.Context) {
	bannerId := c.Param("id")
	id, err := strconv.ParseInt(bannerId, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.DeleteBanner(context.WithValue(context.Background(), "ginContext", c), &proto.BannerRequest{
		Id: int32(id),
	})
	if err != nil {
		api.HandleGrpcErrToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, "")
}

// UpdateBanner 更新轮播图
func UpdateBanner(c *gin.Context) {
	updateBanner := forms.BannerForm{}
	if err := c.ShouldBindJSON(&updateBanner); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}

	fmt.Println("banner:", updateBanner)

	bannerId := c.Param("id")
	id, err := strconv.ParseInt(bannerId, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	_, err = global.GoodsSrvClient.UpdateBanner(context.WithValue(context.Background(), "ginContext", c), &proto.BannerRequest{
		Id:    int32(id),
		Index: int32(updateBanner.Index),
		Image: updateBanner.Image,
		Url:   updateBanner.Url,
	})
	if err != nil {
		api.HandleGrpcErrToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
	})

}
