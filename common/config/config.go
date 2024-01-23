package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// 总配置 抽象 用于对外暴露
type config struct {
	Server server `yaml:"server"`
	Db     db `yaml:"db"`
	Redis  redis `yaml:"redis"`
	Log    log `yaml:"log"`
	ImageSettings imageSettings `yaml:"imageSettings"`
}

// 项目端口 配置抽象
type server struct {
	Address string `yaml:"address"`

	Mode    string `yaml:"mode"`
}

// 数据库配置抽象
type db struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

// redis配置抽象
type redis struct {
	Address string `yaml:"address"`
	Password string `yaml:"password"`
}


// imageSettings图片上传配置
type imageSettings struct {
	UploadDir string `yaml:"uploadDir"`
	ImageHost string `yaml:"imageHost"`
}

// log日志
type log struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

var Config *config

// 配置初始化
func init() {
	file, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		panic(err)
	}
}
