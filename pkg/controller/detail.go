package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArticelDetail(c *gin.Context) {
	c.HTML(http.StatusOK, "detail.html", gin.H{
		"list": []string{"测试Blog标题一", "测试Blog标题二", "测试Blog标题三"},
	})
}
