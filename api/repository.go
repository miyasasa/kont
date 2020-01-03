package api

import (
	"github.com/gin-gonic/gin"
	"kont/internal/remoterepository/bitbucket"
	"kont/internal/repository"
	"kont/storage"
)

func init() {
	group := Router.Group("/repository")
	group.POST("/", saveRepository)
	group.GET("/:name", getRepository)
	group.GET("/", getRepositories)
	group.DELETE("/:name", deleteRepository)
}

func saveRepository(c *gin.Context) {
	var repo = &repository.Repository{}

	_ = c.BindJSON(repo)

	repo.Initialize()
	bitbucket.UpdateUsers(repo)

	err := storage.Storage.PUT(repo.Name, repo)

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
