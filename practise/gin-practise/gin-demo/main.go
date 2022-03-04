package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"kube-learning/practise/gin-practise/gin-demo/log"
	"kube-learning/practise/gin-practise/gin-demo/middleware"
	"kube-learning/practise/gin-practise/gin-demo/response"
	"kube-learning/practise/gin-practise/gin-demo/validator"
)


type options struct {
	addr 	string
	engine  *gin.Engine
}

func NewHttpServer(addr string) *options {

	gin.SetMode(gin.DebugMode)

	gin.DefaultWriter = io.MultiWriter(log.AccessLog())

	validator.Validators()

	o := options{
		addr: addr,
		engine: gin.Default(),
	}

	o.registerHttpRoute()

	return &o
}

func (o *options) registerHttpRoute() {
	o.engine.Use(middleware.LoggerToFile())
	request := o.engine.Group("/practise")
	request.GET("/cookie", response.CookiePractise)
	request.GET("/get", middleware.AuthServer(), response.GetPractise)
	request.POST("/post", response.PostPractise)

}

func (o *options) run() {
	o.engine.Run(o.addr)
}

func main() {
	s := NewHttpServer(":8080")

	s.run()
}