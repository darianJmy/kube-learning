package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("*")
	r.POST("/form", func(c *gin.Context) {
		c.Request.ParseForm()
		data, _ := ioutil.ReadAll(c.Request.Body)
		c.String(http.StatusOK, fmt.Sprintf("ctx.Request.body: %v", string(data)))
	})
	r.GET("/form", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.Run(":8080")

}
