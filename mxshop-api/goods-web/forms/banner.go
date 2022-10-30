package forms

//BannerForm 轮播图表单验证
//type BannerForm struct {
//	Index int    `form:"index" json:"index"  binding:"required"`
//	Image string `form:"image" json:"image" binding:"url"`
//	Url   string `form:"url" json:"url" binding:"url"`
//}

type BannerForm struct {
	Image string `form:"image" json:"image" binding:"url"`
	Index int    `form:"index" json:"index" binding:"required"`
	Url   string `form:"url" json:"url" binding:"url"`
}
