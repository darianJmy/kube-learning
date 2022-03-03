package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/long_async", long_async)
	r.GET("long_sync", long_sync)
	r.Run(":8080")
}

func long_async(c *gin.Context) {
	copyContext := c.Copy()
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("异步执行" + copyContext.Request.URL.Path)
	}()
}

func long_sync(c *gin.Context) {
	time.Sleep(4 * time.Second)
	log.Println("同步执行" + c.Request.URL.Path)
}