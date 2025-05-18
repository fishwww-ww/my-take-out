package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var envPtr = pflag.String("env", "dev", "Environment: dev or prod")

func InitLoadConfig() *AllConfig {
	pflag.Parse()

	config := viper.New()
	// 设置读取路径
	config.AddConfigPath("./config")
	// 设置读取文件名字
	config.SetConfigName(fmt.Sprintf("application-%s", *envPtr))
	// 设置读取文件类型
	config.SetConfigType("yaml")
	// 读取文件载体
	var configData *AllConfig
	// 读取配置文件
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Use Viper ReadInConfig Fatal error config err:%s \n", err))
	}
	// 查找对应配置文件
	err = config.Unmarshal(&configData)
	if err != nil {
		panic(fmt.Errorf("read config file to struct err: %s\n", err))
	}
	return configData
}

type AllConfig struct {
	Server     Server
	DataSource DataSource
	Jwt        Jwt
	Log        Log
	Redis      Redis
	AliOss     AliOss
}

type Server struct {
	Port  string
	Level string
}
type DataSource struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"` // 用db_name会报错,很奇怪
	Config   string `yaml:"config"`
}

type Jwt struct {
	Admin JwtOption
	User  JwtOption
}

type Log struct {
	Level    string
	FilePath string
}

type Redis struct {
	Host     string
	Port     string
	Password string
	DataBase int `mapstructure:"data_base"`
}

type JwtOption struct {
	Secret string
	TTL    string
	Name   string
}

type AliOss struct {
	EndPoint        string
	AccessKeyId     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
	BucketName      string `mapstructure:"bucket_name"`
}

func (d *DataSource) Dsn() string {
	return d.Username + ":" + d.Password + "@tcp(" + d.Host + ":" + d.Port + ")/" + d.DBName + "?" + d.Config
}
