package forms

//ShopCartForm 购物车表单
type ShopCartForm struct {
	GoodsId int32 `json:"goods_id" form:"goods_id" binding:"required"`
	Nums    int32 `json:"nums" form:"nums" binding:"required,min=1"`
}

//UpdateShopCartForm 更新购物车状态
type UpdateShopCartForm struct {
	Nums    int32 `json:"nums" form:"nums" binding:"required,min=1"`
	Checked *bool `json:"checked" form:"checked"`
}
