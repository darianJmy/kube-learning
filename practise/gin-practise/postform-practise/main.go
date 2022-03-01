package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("index.html")

	r.POST("/from", func(c *gin.Context) {
		name := c.PostForm("username")
		userpassword := c.PostForm("userpassword")
		c.String(http.StatusOK, name + userpassword)
	})
	r.GET("/from", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.Run(":8080")

}
