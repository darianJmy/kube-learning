package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 当发起这个请求时应当对另一个 api 发起网络请求，并且返回给客户端
	r.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.JSON(http.StatusServiceUnavailable, gin.H{"message": "获取数据请求失败", "status": http.StatusServiceUnavailable})
			return
		}
		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Contest-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	r.Run(":8080")
}
