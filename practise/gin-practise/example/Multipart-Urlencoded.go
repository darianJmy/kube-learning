package main

import (
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var form LoginForm
		// 显式绑定，就是指定结构体于数据来源
		// c.ShouldBindWith(&form, binding.Form)
		// 隐式绑定，使用简单，但是隐式绑定也是调用的显式绑定
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(401, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(200, gin.H{"status": "you are logged in"})
	})
	router.Run(":8080")
}
