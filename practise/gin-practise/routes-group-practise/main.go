package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("/logout", logout)
	}
	v2 := r.Group("/v2")
	v2.POST("/postname", postname)
	v2.POST("/postpasswd", postpasswd)

	r.Run(":8080")
}

func login(c *gin.Context) {
	c.String(http.StatusOK,"login")
}

func logout(c *gin.Context) {
	c.String(http.StatusOK, "logout")
}

func postname(c *gin.Context) {
	body := c.Request.Body
	x, _ := io.ReadAll(body)
	c.JSON(http.StatusOK, x)
}

func postpasswd(c *gin.Context) {

}