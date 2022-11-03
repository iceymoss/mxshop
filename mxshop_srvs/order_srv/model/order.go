package model

import "time"

type ShoppingCart struct {
	BaseModel
	User    int32 `gorm:"type:int;index"` //用户索引，用来快速当前用户的购物车记录
	Goods   int32 `gorm:"type:int;index"` //加索引：需要查询的时候
	Nums    int32 `gorm:"type:int"`       //商品数量
	Checked bool  //是否选中
}

//TableName 表名
func (ShoppingCart) TableName() string {
	return "shoppingcart"
}

//OrderInfo 订单信息
type OrderInfo struct {
	BaseModel

	User    int32  `gorm:"type:int;index"`
	OrderSn string `gorm:"type:varchar(30);index"` //订单号，平台自己生成的订单号
	PayType string `gorm:"type:varchar(20) comment 'alipay(支付宝)， wechat(微信)'"`

	//status大家可以考虑使用iota来做
	Status     string     `gorm:"type:varchar(20)  comment 'PAYING(待支付), TRADE_SUCCESS(成功)， TRADE_CLOSED(超时关闭), WAIT_BUYER_PAY(交易创建), TRADE_FINISHED(交易结束)'"`
	TradeNo    string     `gorm:"type:varchar(100) comment '交易号'"` //交易号，其实就是支付宝或者微信的的订单号，用于查账
	OrderMount float32    //订单金额
	PayTime    *time.Time `gorm:"type:datetime"`

	Address      string `gorm:"type:varchar(100)"`
	SignerName   string `gorm:"type:varchar(20)"`
	SingerMobile string `gorm:"type:varchar(11)"`
	Post         string `gorm:"type:varchar(20)"` //留言信息
}

func (OrderInfo) TableName() string {
	return "orderinfo"
}

//OrderGoods 订单商品信息
type OrderGoods struct {
	BaseModel

	Order int32 `gorm:"type:int;index"`
	Goods int32 `gorm:"type:int;index"`

	//把商品信息保存下来，但是字段冗余，不符合mysql三范式，但是高并发系统一般都不遵循三范式
	GoodsName  string `gorm:"type:varchar(100);index"`
	GoodsImage string `gorm:"type:varchar(200)"`
	GoodsPrice float32
	Nums       int32 `gorm:"type:int"`
}

func (OrderGoods) TableName() string {
	return "ordergoods"
}
