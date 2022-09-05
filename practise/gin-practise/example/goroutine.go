package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	r := gin.Default()

	// 记录到文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 协程
	r.GET("/long_async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path" + cCp.Request.URL.Path)
		}()
	})
	// 单线程
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path" + c.Request.URL.Path)
	})

	r.Run(":8080")
}
