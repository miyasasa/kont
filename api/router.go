package api

import (
	"github.com/gin-gonic/gin"
	"miya/api/ping"
)

func InitRouter() *gin.Engine {

	gn := gin.Default()

	gn.GET("/ping", ping.Ping)

	return gn
}
