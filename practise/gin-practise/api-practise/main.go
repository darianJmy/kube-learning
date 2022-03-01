package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	router := gin.Default()

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name +" is "+ action)
	})
	router.Run(":8081")
}
