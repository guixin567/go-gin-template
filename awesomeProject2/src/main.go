package main

import (
	"awesomeProject/awesomeProject2/src/config"
	"awesomeProject/awesomeProject2/src/dao/mysql"
	"awesomeProject/awesomeProject2/src/dao/redis"
	"awesomeProject/awesomeProject2/src/log"
	"awesomeProject/awesomeProject2/src/router"
	"awesomeProject/awesomeProject2/src/router/test"
	"fmt"
	"go.uber.org/zap"
)

func main() {

	//1 初始化配置
	if err := config.Init(); err != nil {
		fmt.Println("config error", err)
	}

	//2 初始化日志
	defer log.ALog.Sync()
	if err := log.Init(); err != nil {
		fmt.Println("log error", err)
	}

	//3 初始化mysql
	if err := mysql.Init(); err != nil {
		log.ALog.Error("mysql init error", zap.Error(err))
	}
	defer mysql.Xdb.Close()

	//4 初始化redis
	if err := redis.Init(); err != nil {
		log.ALog.Error("redis init error", zap.Error(err))
	}
	defer redis.Rdb.Close()

	//5 初始化路由
	app := router.Register(test.Router)

	//6 开启服务
	if err := app.Run(fmt.Sprintf(":s%", config.GlobalConfig.App.Port)); err != nil {
		log.ALog.Error("router init error", zap.Error(err))
	}
}
