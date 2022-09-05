package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/login", login)
	r.GET("/home", authmiddleware(), home)
	r.Run(":8080")

}

func authmiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("abc"); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
			c.Abort()
		} else if cookie != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "cookie value is error"})
			c.Abort()
		}
	}
}

func login(c *gin.Context) {
	cookie := &http.Cookie{
		Name:     "abc",
		Value:    "123",
		Path:     "/",
		MaxAge:   60,
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "home"})
}
