package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"miya/internal/repository"
	"miya/storage"
)

func init() {
	group := Router.Group("/repository")
	group.POST("/", saveRepository)
	group.GET("/", getRepositories)
	group.DELETE("/:name", deleteRepository)
}

func saveRepository(c *gin.Context) {
	var repo repository.Repository

	err := c.BindJSON(&repo)
	if err != nil {
		log.Printf("repository::saveRepository bind exception")
	}

	// get provider
	// init users

	err = storage.Storage.PUT(repo.Name, repo)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, repo)
}

func deleteRepository(c *gin.Context) {
	name := c.Param("name")

	err := storage.Storage.Delete(name)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(204, gin.H{
		"message": "Repository deleted ...",
	})

}

func getRepositories(c *gin.Context) {
	repos := storage.Storage.GetAllRepositories()
	c.JSON(200, repos)
}
