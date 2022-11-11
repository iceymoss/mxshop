package router

import (
	"mxshop-api/goods-web/api/brands"
	"mxshop-api/goods-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitBrands(router *gin.RouterGroup) {
	BrandsRouter := router.Group("brands").Use(middlewares.Trace())
	{
		BrandsRouter.GET("", brands.BrandList)                                                        //获取品牌列表
		BrandsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdmin(), brands.NewBrand)          //添加品牌
		BrandsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), brands.DeleteBrand) //删除品牌
		BrandsRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), brands.UpdateBrand)    //修改品牌
	}

	CategoryBrandRouter := router.Group("categorybrands").Use(middlewares.Trace())
	{
		CategoryBrandRouter.GET("", brands.CategoryBrandList)                                                        // 类别品牌列表页
		CategoryBrandRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), brands.DeleteCategoryBrand) // 删除类别品牌
		CategoryBrandRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdmin(), brands.NewCategoryBrand)          //新建类别品牌
		CategoryBrandRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), brands.UpdateCategoryBrand)    //修改类别品牌
		CategoryBrandRouter.GET("/:id", brands.GetCategoryBrandList)                                                 //获取分类的品牌
	}
}
