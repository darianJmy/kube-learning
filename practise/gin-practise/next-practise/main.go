package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(middleware())
	r.GET("/request", func(c *gin.Context){
		req, _ := c.Get("request")
		fmt.Println("request", req)
		c.JSON(http.StatusOK, gin.H{
			"request": req,
		})
	})
	r.Run(":8080")
}


func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		c.Set("request", "中间件")
		c.Next()

		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)

		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}