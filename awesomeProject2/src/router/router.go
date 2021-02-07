package router

import "github.com/gin-gonic/gin"

type Option func(engin *gin.Engine)

var Options = []Option{}

//注册路由
func Register(opts ...Option) *gin.Engine {
	Options = append(Options, opts...)
	return Init()
}

//初始化路由
func Init() *gin.Engine {
	engine := gin.Default()
	for _, Opt := range Options {
		Opt(engine)
	}

	return engine

}
