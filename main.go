package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/likeyao/peach-blog/pkg/controller"
)

func main() {
	var path string

	flag.StringVar(&path, "p", "/Users/likeyao/demo/blogs", "存放博客文章的文件夹")
	flag.Parse()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.Index(path))
	r.GET("/article/detail", controller.ArticelDetail(path))
	r.Static("/assets", "assets")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//要能够根据配置，去某个文件目录中读取所有的博客内容，文件名作为标题、读取文件中的一段内容作为预览内容
