package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}
