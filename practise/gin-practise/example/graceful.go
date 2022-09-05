package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// 优雅的关机，意思就是关闭进程时不接受新的请求会把正在处理的请求处理完再退出。

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	// 协程
	go func() {
		// 启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s \n", err)
		}
	}()

	// 初始化 chan，接收 os.Signal
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer 闭包
	defer cancel()
	// 主动关闭 http 服务
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
