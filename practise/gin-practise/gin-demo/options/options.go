package options

import (
	"github.com/gin-gonic/gin"
	"kube-learning/practise/gin-practise/gin-demo/middleware"
	"kube-learning/practise/gin-practise/gin-demo/response"
)


type Options struct {
	Addr   string
	Engine *gin.Engine
}


func (o *Options) RegisterHttpRoute() {
	o.Engine.Use(middleware.LoggerToFile())
	request := o.Engine.Group("/practise")
	request.GET("/cookie", response.CookiePractise)
	request.GET("/get", middleware.AuthServer(), response.GetPractise)
	request.POST("/post", response.PostPractise)
	request.GET("/async", response.AsyncPractise)
	request.POST("/upload", response.UploadPractise)

}

func (o *Options) Run() {
	o.Engine.Run(o.Addr)
}
