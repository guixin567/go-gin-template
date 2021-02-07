package test

import "github.com/gin-gonic/gin"

func Router(engin *gin.Engine) {
	engin.GET("/test", TestHandler)
}
