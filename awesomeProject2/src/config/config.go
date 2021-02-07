package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var GlobalConfig *AppConfig

func Init() (err error) {
	viper.SetConfigFile("./src/config/config.json")
	err = viper.ReadInConfig()
	GlobalConfig = new(AppConfig)
	err = viper.Unmarshal(GlobalConfig)
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		println("config change")
	})
	return
}

type AppConfig struct {
	*App   `mapstructure:"app"`
	*Redis `mapstructure:"redis"`
	*MySql `mapstructure:"mysql"`
}

type App struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type MySql struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   int    `mapstructure:"dbname"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}
