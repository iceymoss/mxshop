package config

//GoodsSerConfig 映射商品配置
type GoodsSerConfig struct {
	Name string `mapstructure:"name" json:"name"`
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

//JWTConfig 映射token配置
type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

//AliSmsConfig 阿里秘钥
type AliSmsConfig struct {
	Apikey    string `mapstructure:"key" json:"key"`
	ApiSecret string `mapstructure:"secret" json:"secret"`
}

//ConsulConfig 注册中心配置
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

//ServerConfig  映射服务配置
type ServerConfig struct {
	Name             string         `mapstructure:"name" json:"name"`
	Host             string         `mapstructure:"host" json:"host"`
	Port             int            `mapstructure:"port" json:"port"`
	Tag              []string       `mapstructure:"tag" json:"tag"`
	GoodsSerInfo     GoodsSerConfig `mapstructure:"goods_srv" json:"goods_srv"`
	OrderSerInfo     GoodsSerConfig `mapstructure:"order_srv" json:"order_srv"`
	InventorySerInfo GoodsSerConfig `mapstructure:"inventory_srv" json:"inventory_srv"`
	JWTInfo          JWTConfig      `mapstructure:"jwt" json:"jwt"`
	AliSms           AliSmsConfig   `mapstructure:"sms" json:"sms"`
	ConsulInfo       ConsulConfig   `mapstructure:"consul" json:"consul"`
}

//NacosConfig 配置中心配置
type NacosConfig struct {
	Host        string `mapstructure:"host" json:"host"`
	Port        uint64 `mapstructure:"port" json:"port"`
	NamespaceId string `mapstructure:"namespace_id" json:"namespace_id"`
	User        string `mapstructure:"user" json:"user"`
	Password    string `mapstructure:"password" json:"password"`
	DataId      string `mapstructure:"data_id" json:"data_id"`
	Group       string `mapstructure:"group" json:"group"`
}
