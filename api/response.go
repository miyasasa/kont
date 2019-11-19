package api

import "github.com/gin-gonic/gin"

func okOrElse404(err error, c *gin.Context, i interface{}) {
	response(err, c, 200, i, 404)
}

func okOrElse500(err error, c *gin.Context, i interface{}) {
	response(err, c, 200, i, 500)

}

func noContentOrElse404(err error, c *gin.Context) {
	response(err, c, 204, nil, 404)
}

func ok(c *gin.Context, i interface{}) {
	c.JSON(200, i)
}

func response(err error, c *gin.Context, c1 int, i interface{}, c2 int) {
	if err != nil {
		c.JSON(c2, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(c1, i)
	}
}
