package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type User struct {
	Age     int    `json:"age" form:"age" binding:"required,gt=10"`
	Name    string `json:"name" form:"name" binding:"NotNullAndAdmin"`
	Address string `json:"address" form:"address" binding:"required"`
}

func main() {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 这里的 key 和 fn 可以不一样最终在 struct 使用的是 key
		v.RegisterValidation("NotNullAndAdmin", nameNotNullAndAdmin)
	}

	r.POST("/login", login)
	r.Run(":8080")
}

func nameNotNullAndAdmin(fl validator.FieldLevel) bool {

	value := fl.Field().String()
	return value != "" && "admin" != value
}

func login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	c.JSON(http.StatusOK, user)
}
