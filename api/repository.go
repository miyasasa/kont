package api

import (
	"github.com/gin-gonic/gin"
)

func init() {
	group := Router.Group("/repository")
	group.POST("/", saveRepository)
	group.GET("/", getRepositories)
}

func saveRepository(c *gin.Context) {
	c.JSON(200, gin.H{
		"Repository": "saved ....",
	})
}

func getRepositories(c *gin.Context) {
	c.JSON(200, gin.H{
		"repositories": "repos are coming.....",
	})
}
