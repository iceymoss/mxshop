package model

//LeavingMessages 用户留言
type LeavingMessages struct {
	BaseModel
	User        int32  `gorm:"type:int;index"`
	MessageType int32  `gorm:"type:int comment '留言类型: 1(留言),2(投诉),3(询问),4(售后),5(求购)'"`
	Subject     string `gorm:"type:varchar(100)"`

	Message string
	File    string `gorm:"type:varchar(200)"` //上传图片url
}

func (LeavingMessages) TableName() string {
	return "leavingmessages"
}

//Address 收货信息
type Address struct {
	BaseModel

	User         int32  `gorm:"type:int;index"`
	Province     string `gorm:"type:varchar(10)"`
	City         string `gorm:"type:varchar(10)"`
	District     string `gorm:"type:varchar(20)"`
	Address      string `gorm:"type:varchar(100)"`
	SignerName   string `gorm:"type:varchar(20)"`
	SignerMobile string `gorm:"type:varchar(11)"`
}

//UserFav 用户收藏
type UserFav struct {
	BaseModel

	//用户不能多次收藏相同商品
	User  int32 `gorm:"type:int;index:idx_user_goods,unique"`
	Goods int32 `gorm:"type:int;index:idx_user_goods,unique"`
}

func (UserFav) TableName() string {
	return "userfav"
}
