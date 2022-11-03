package config

//MysqlConfig mysql信息配置
type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

//ConsulConfig consul配置
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type RedisConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

//GoodsSerConfig 映射商品配置
type GoodsSerConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

//MqConfig 消息队列配置
type MqConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	GroupName string `mapstructure:"group_name" json:"group_name"`
	Topic     string `mapstructure:"topic" json:"topic"`
}

//ServerConfig 服务配置
type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name"`
	Host       string       `mapstructure:"host" json:"host"`
	Port       int          `mapstructure:"port" json:"port"`
	MysqlInfo  MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
	RedisInfo  RedisConfig  `mapstructure:"redis" json:"redis"`
	Tags       []string     `mapstructure:"tags" json:"tags"`

	//商品微服务
	GoodsSerInfo GoodsSerConfig `mapstructure:"goods_srv" json:"goods_srv"`
	//库存微服务
	InventorySerInfo GoodsSerConfig `mapstructure:"inventory_srv" json:"inventory_srv"`
	MqInfo           MqConfig       `mapstructure:"rocket" json:"rocket"`
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

//NacosServer 配置中心
type NacosServer struct {
	NacosInfo NacosConfig `mapstructure:"nacos" json:"nacos"`
}
