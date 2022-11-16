package config

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name"`
	Host       string       `mapstructure:"host" json:"host"`
	Tags       []string     `mapstructure:"tags" json:"tags"`
	Port       int          `mapstructure:"port" json:"port"`
	JWTInfo    JWTConfig    `mapstructure:"jwt" json:"jwt"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
	OssInfo    OssConfig    `mapstructure:"oss" json:"oss"`
}

type OssConfig struct {
	ApiKey      string `mapstructure:"key" json:"key"`
	ApiSecrect  string `mapstructure:"secrect" json:"secrect"`
	Host        string `mapstructure:"host" json:"host"`
	CallBackUrl string `mapstructure:"callback_url" json:"callback_url"`
	UploadDir   string `mapstructure:"upload_dir" json:"upload_dir"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}
