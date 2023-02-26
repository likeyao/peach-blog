package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", index)
	r.GET("/ping", pong)
	r.Static("/assets", "assets")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"list": []string{"测试Blog标题一", "测试Blog标题二", "测试Blog标题三"},
	})
}
