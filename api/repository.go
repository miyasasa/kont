package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"miya/internal/repository"
)

func init() {
	group := Router.Group("/repository")
	group.POST("/", saveRepository)
	group.GET("/", getRepositories)
}

func saveRepository(c *gin.Context) {
	var repo repository.Repository

	err := c.BindJSON(&repo)
	if err != nil {
		log.Printf("repository::saveRepository bind exception")
	}

	// get provider
	// init users

	c.JSON(200, repo)
}

func getRepositories(c *gin.Context) {
	c.JSON(200, gin.H{
		"repositories": "repos are coming.....",
	})
}
