package controller

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	Path string
}

func (i *IndexController) Index(c *gin.Context) {
	//试一下读取文件名
	articles := readArticles(i.Path)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"list": articles,
	})
}

func readArticles(path string) []article {
	result := make([]article, 0)
	dirs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, v := range dirs {
		fullPath := path + "/" + v.Name()
		if !v.IsDir() {
			result = append(result, article{
				Name:    parseAritcleTitle(v.Name()),
				Content: readArticle(fullPath),
				Outline: readArticle(fullPath),
				Time:    parseTime(v),
			})
		}

	}
	return result
}

func parseAritcleTitle(filename string) string {
	return filename[:strings.Index(filename, ".")]
}

func readArticle(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func parseTime(dir os.DirEntry) string {
	info, err := dir.Info()
	if err != nil {
		panic(err)
	}
	return info.ModTime().Format("2006-01-02 15:04:05")
}

type article struct {
	Name    string
	Content string
	Outline string
	Time    string
}
