package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 使用 LoadHTMLGlob() 或者 LoadHTMLFiles() 加载 html 文件，只要 get index 就会打开静态页面
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}
