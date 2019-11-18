package repository

import "github.com/gin-gonic/gin"

func InitRepository(router *gin.Engine) {
	group := router.Group("/repository")
	group.GET("/", getRepositories)

}

func getRepositories(c *gin.Context) {
	c.JSON(200, gin.H{
		"repositories": "repos are coming.....",
	})
}
