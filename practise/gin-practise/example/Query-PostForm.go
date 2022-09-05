package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {
		// Query 是参数，就是 /post?id=1
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		// PostForm 就是 body 部分响应类型，有 row、none、form-data 等
		// PostForm 就是接收 form-data
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	router.Run(":8080")
}
