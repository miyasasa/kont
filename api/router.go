package api

import (
	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"miya/api/ping"
	"miya/api/repository"
	"net/http"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.HTMLRender = gintemplate.Default()
	router.Static("/assets", "./assets")

	router.GET("/ping", ping.Ping)

	repository.InitRepositories(router)

	// pages dummy APIs
	router.GET("/", func(ctx *gin.Context) {
		//render with master
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	router.GET("/page", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Page file title!!"})
	})

	return router
}
