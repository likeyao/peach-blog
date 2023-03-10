package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArticelDetail(path string) func(c *gin.Context) {
	return func(c *gin.Context) {

		title := c.Query("title")

		c.HTML(http.StatusOK, "detail.html", gin.H{
			"article": readArticle(path, title+".md"),
		})
	}
}
