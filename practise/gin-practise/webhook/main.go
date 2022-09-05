package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	"kube-learning/practise/gin-practise/gin-demo/validator"
)

func PostPractise(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	c.String(http.StatusOK, "OK")
}
func GetPractise(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	action = strings.Trim(action, "/")
	c.String(http.StatusOK, name+" is "+action)
}

func main() {

	validator.Validators()
	router := gin.Default()
	router.POST("/webhook", PostPractise)
	router.GET("/test", GetPractise)
	router.Run(":5001")
}
