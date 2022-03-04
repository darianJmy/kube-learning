package middleware

import (
	"github.com/gin-gonic/gin"
	"kube-learning/practise/gin-practise/gin-demo/log"
	"kube-learning/practise/gin-practise/gin-demo/response"
	"time"
)

func LoggerToFile() gin.HandlerFunc {
	handlerFunc := func(c *gin.Context) {
		startTime := time.Now()

		// 处理请求操作
		c.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		log.Info.Printf("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIp, reqMethod, reqUri)
	}
	return handlerFunc

}

func AuthServer() gin.HandlerFunc {
	r := response.GinResp{}
	handlerFunc := func(c *gin.Context) {
		cookie, err := c.Cookie("abc")
		if err != nil && cookie != "123" {
			r.SetCode(400)
			r.SetMessage("not cookie")
			r.SetError(err)
			c.JSON(400, r)
			c.Abort()
		}
	}
	return handlerFunc
}