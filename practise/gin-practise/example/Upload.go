package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 此处使用 default 是因为启用两个中间件 Logger 和 Recovery 中间件
	// 如果不想使用中间件可以 gin.New()
	r := gin.Default()

	r.POST("/SingleUpload", func(c *gin.Context) {
		// 单文件上传使用 FormFile 注意返回的结构体 multipart.FileHeader
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		dst := "./" + file.Filename
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "上传文件成功", "status": http.StatusOK})
	})

	r.POST("/MultipleUpload", func(c *gin.Context) {
		// 多文件上传使用 MultipartForm 返回的结构体是 multipart.Form
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		// 此时 form.File 返回的结构体是 map[string][]*FileHeader
		// 一个字典，k是 string， v 数组，数组里面内容是 *FileHeader
		// 所以要获取字典里的 k 下的 v ，然后在循环 v 里面内容

		files := form.File["upload[]"]

		for _, file := range files {
			dst := "./" + file.Filename
			log.Println(file.Filename)

			if err := c.SaveUploadedFile(file, dst); err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": "上传文件成功", "status": http.StatusOK})
	})
	r.Run(":8080")
}
