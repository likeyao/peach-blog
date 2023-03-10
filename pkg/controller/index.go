package controller

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Index(path string) func(c *gin.Context) {
	return func(c *gin.Context) {
		articles := readArticles(path)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"list": articles,
		})
	}
}

func readArticles(path string) []article {
	result := make([]article, 0)
	dirs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, v := range dirs {
		if !v.IsDir() {
			result = append(result, readArticle(path, v.Name()))
		}

	}
	return result
}

func readArticle(path string, filename string) article {
	fullPath := path + "/" + filename

	file, err := os.Open(fullPath)
	if err != nil {
		panic(err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	return article{
		Name:    parseAritcleTitle(filename),
		Content: content(fullPath),
		Outline: content(fullPath),
		Time:    getModifyTime(fileInfo),
	}
}

func parseAritcleTitle(filename string) string {
	return filename[:strings.Index(filename, ".")]
}

func content(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func getModifyTime(info os.FileInfo) string {
	return info.ModTime().Format("2006-01-02 15:04:05")
}

type article struct {
	Name    string
	Content string
	Outline string
	Time    string
}
