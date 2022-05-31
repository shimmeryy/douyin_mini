package config

import (
	"github.com/spf13/viper"
	"log"
)

var AppConfig *Config

func init() {
	var ViperConfig = viper.New()
	InitConfig(ViperConfig)
	if err := ViperConfig.Unmarshal(&AppConfig); err != nil {
		log.Printf("初始话配置文件失败！")
	}
}

func InitConfig(viperConfig *viper.Viper) {
	viperConfig.AddConfigPath("./")          //设置读取的文件路径
	viperConfig.SetConfigName("application") //设置读取的文件名
	viperConfig.SetConfigType("yaml")        //设置文件的类型
	//读取配置
	if err := viperConfig.ReadInConfig(); err != nil {
		//预防测试路径时读取配置文件
		viperConfig.AddConfigPath("../../../")
		if err := viperConfig.ReadInConfig(); err != nil {
			panic(err)
		}
	}
}

type Config struct {
	Server     Server
	DataSource DataSource
	Redis      Redis
}

type Server struct {
	Port string `mapstructure:"port"`
}

type DataSource struct {
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}
