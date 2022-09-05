package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type User struct {
	Name     string    `json:"name" form:"name" binding:"required"`
	Age      int       `json:"age" form:"age" binding:"required,gt=10"`
	Birthday time.Time `json:"birthday" time_form:"2006-01-02" time_utc:"1"`
}

func main() {
	r := gin.Default()
	r.POST("/login", login)
	r.Run(":8080")
}

func login(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.String(http.StatusOK, fmt.Sprintf("%#v", user))
}
