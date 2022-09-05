package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}
type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func main() {
	r := gin.Default()
	// 多次绑定 ShouldBindBodyWith 可以根据数据类型来操作
	r.POST("/someJSON", func(c *gin.Context) {
		objA := formA{}
		objB := formB{}
		if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
			c.String(200, "the body should be formA")
		} else if errB := c.ShouldBindBodyWith(&objB, binding.XML); errB == nil {
			c.String(200, "the body should be formB")
		}
	})
	r.Run(":8080")
}
