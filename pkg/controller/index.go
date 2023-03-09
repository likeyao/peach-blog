package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"list": []string{"测试Blog标题一", "测试Blog标题二", "测试Blog标题三"},
	})
}
