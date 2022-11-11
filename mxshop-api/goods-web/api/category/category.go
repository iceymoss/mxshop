package category

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"mxshop-api/goods-web/api"
	"mxshop-api/goods-web/forms"
	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/proto"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
)

func List(c *gin.Context) {
	Rsp, err := global.GoodsSrvClient.GetAllCategorysList(context.WithValue(context.Background(), "ginContext", c), &empty.Empty{})
	if err != nil {
		api.HandleGrpcErrToHttp(err, c)
		return
	}
	data := make([]interface{}, 0)

	//json反序列化
	err = json.Unmarshal([]byte(Rsp.JsonData), &data)
	if err != nil {
		zap.S().Info("List [查询] 【分类列表】失败", err.Error())
	}
	c.JSON(http.StatusOK, data)
}

func Detail(c *gin.Context) {
	goodsId := c.Param("id")
	id, err := strconv.ParseInt(goodsId, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	reMap := make(map[string]interface{})
	subCategorys := make([]interface{}, 0)
	Rsp, err := global.GoodsSrvClient.GetSubCategory(context.WithValue(context.Background(), "ginContext", c), &proto.CategoryListRequest{
		Id: int32(id),
	})
	if err != nil {
		api.HandleGrpcErrToHttp(err, c)
		return
	} else {
		for _, value := range Rsp.SubCategorys {
			subCategorys = append(subCategorys, map[string]interface{}{
				"id":              value.Id,
				"name":            value.Name,
				"level":           value.Level,
				"parent_category": value.ParentCategory,
				"is_tab":          value.IsTab,
			})
		}
		reMap["id"] = Rsp.Info.Id
		reMap["name"] = Rsp.Info.Name
		reMap["level"] = Rsp.Info.Level
		reMap["parent_category"] = Rsp.Info.ParentCategory
		reMap["is_tab"] = Rsp.Info.IsTab
		reMap["sub_categorys"] = Rsp.SubCategorys

		c.JSON(http.StatusOK, reMap)
	}
	return
}

func New(c *gin.Context) {
	var categoryform forms.CategoryForm
	if err := c.ShouldBindJSON(&categoryform); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}

	Rsp, err := global.GoodsSrvClient.CreateCategory(context.WithValue(context.Background(), "ginContext", c), &proto.CategoryInfoRequest{
		Name:           categoryform.Name,
		ParentCategory: categoryform.ParentCategory,
		Level:          categoryform.Level,
		IsTab:          *categoryform.IsTab,
	})
	if err != nil {
		api.HandleGrpcErrToHttp(err, c)
		return
	}

	request := make(map[string]interface{})
	request["id"] = Rsp.Id
	request["name"] = Rsp.Name
	request["parent"] = Rsp.ParentCategory
	request["level"] = Rsp.Level
	request["is_tab"] = Rsp.IsTab

	c.JSON(http.StatusOK, request)
}

func Delete(c *gin.Context) {
	categoryId := c.Param("id")
	id, err := strconv.ParseInt(categoryId, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	_, err = global.GoodsSrvClient.DeleteCategory(context.WithValue(context.Background(), "ginContext", c), &proto.DeleteCategoryRequest{
		Id: int32(id),
	})
	if err != nil {
		api.HandleGrpcErrToHttp(err, c)
		return
	}
	c.Status(http.StatusOK)
}

func Update(c *gin.Context) {
	updatecategoryform := forms.UpdateCategoryForm{}
	if err := c.ShouldBindJSON(&updatecategoryform); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}

	goodsId := c.Param("id")
	id, err := strconv.ParseInt(goodsId, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	var categoryinfoRequest proto.CategoryInfoRequest

	categoryinfoRequest.Id = int32(id)
	categoryinfoRequest.Name = updatecategoryform.Name

	if updatecategoryform.IsTab != nil {
		categoryinfoRequest.IsTab = *updatecategoryform.IsTab
	}

	_, err = global.GoodsSrvClient.UpdateCategory(context.WithValue(context.Background(), "ginContext", c), &categoryinfoRequest)
	if err != nil {
		api.HandleGrpcErrToHttp(err, c)
		return
	}

	c.Status(http.StatusOK)
}
