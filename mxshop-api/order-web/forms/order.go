package forms

type OrderForms struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	Mobile  string `json:"mobile" form:"mobile" binding:"required,mobile"`
	Post    string `json:"post" form:"post" binding:"required"`
}
