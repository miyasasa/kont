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
	group.GET("/:name", getRepository)
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

	okOrElse500(err, c, repo)
}

func getRepository(c *gin.Context) {
	name := c.Param("name")

	var repo repository.Repository
	err := storage.Storage.GET(name, &repo)

	okOrElse404(err, c, repo)
}

func deleteRepository(c *gin.Context) {
	name := c.Param("name")

	err := storage.Storage.Delete(name)

	noContentOrElse404(err, c)
}

func getRepositories(c *gin.Context) {
	repos := storage.Storage.GetAllRepositories()
	ok(c, repos)
}
