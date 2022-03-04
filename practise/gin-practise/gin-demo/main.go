package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"kube-learning/practise/gin-practise/gin-demo/log"
	"kube-learning/practise/gin-practise/gin-demo/options"
	"kube-learning/practise/gin-practise/gin-demo/validator"
)

func NewHttpServer(addr string) *options.Options {

	gin.SetMode(gin.DebugMode)

	gin.DefaultWriter = io.MultiWriter(log.AccessLog())

	validator.Validators()

	o := options.Options{
		Addr:   addr,
		Engine: gin.Default(),
	}
	o.RegisterHttpRoute()
	return &o
}


func main() {
	s := NewHttpServer(":8080")

	s.Run()
}