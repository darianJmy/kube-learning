package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.POST("upload", upload)
	r.GET("upload", get)
	r.Run(":8080")
}

func upload(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Println(name)
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		panic(err)
	}
	filename := header.Filename
	fmt.Println(file, filename)

	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		panic(err)
	}
	c.String(http.StatusOK, "update successful")
}

func get(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{})
}
