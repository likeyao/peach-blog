package main

import (
	"github.com/gin-gonic/gin"
	"github.com/likeyao/peach-blog/pkg/controller"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.Index)
	r.GET("/article/detail", controller.ArticelDetail)
	r.Static("/assets", "assets")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
