package config

//UserSerConfig 映射用户配置
type UserSerConfig struct {
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

//ParamsConfig 短信模板配置
type ParamsConfig struct {
	SignName     string `mapstructure:"sign_name" json:"sign_name"`
	TemplateCode string `mapstructure:"code" json:"code"`
}

//RedisConfig redis数据库配置
type RedisConfig struct {
	Host  string `mapstructure:"host" json:"host"`
	Port  int    `mapstructure:"port" json:"port"`
	Expir int    `mapstructure:"expir" json:"expir"`
}

//Verifier 手机验证长度
type Verifier struct {
	Width int `mapstructure:"width" json:"width"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

//ServerConfig  映射服务配置
type ServerConfig struct {
	Name        string        `mapstructure:"name" json:"name"`
	Host        string        `mapstructure:"host" json:"host"`
	Port        int           `mapstructure:"port" json:"port"`
	Tag         []string      `mapstructure:"tag" json:"tag"`
	UserSerInfo UserSerConfig `mapstructure:"user_srv" json:"user_srv"`
	JWTInfo     JWTConfig     `mapstructure:"jwt" json:"jwt"`
	AliSms      AliSmsConfig  `mapstructure:"sms" json:"sms"`
	Params      ParamsConfig  `mapstructure:"params" json:"params"`
	Redis       RedisConfig   `mapstructure:"redis" json:"redis"`
	Verify      Verifier      `mapstructure:"verify" json:"verify"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul" json:"consul"`
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
