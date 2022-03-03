package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/login", login)
	v1.GET("/logout", logout)


	v2 := r.Group("/v2")
	v2.POST("/postname", postname)
	v2.POST("/postpasswd", postpasswd)

	r.Run(":8080")
}

func login(c *gin.Context) {
	c.String(http.StatusOK, "login")
}

func logout(c *gin.Context) {
	c.String(http.StatusOK, "logout")
}

func postname(c *gin.Context) {
	name := c.DefaultPostForm("name", "jixingxing")

	c.String(http.StatusOK, fmt.Sprintf("name:%s", name))

}

func postpasswd(c *gin.Context) {
	passwd := c.DefaultPostForm("passwd", "11111")
	c.String(http.StatusOK, fmt.Sprintf("passwd:%s", passwd))
}

// curl http://127.0.0.1:8080/v2/postname -X POST