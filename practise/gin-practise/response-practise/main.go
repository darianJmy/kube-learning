package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name" xml:"name" yaml:"name" uri:"name"`
	Password string `json:"password" xml:"password" yaml:"password" uri:"password"`
	Age int `json:"age" xml:"age" yaml:"age" uri:"age"`
}

//type Message struct {
//	Name xml.Name `json:"name" xml:"name" yaml:"name" uri:"name"`
//	Password xml.Name `json:"password" xml:"password" yaml:"password" uri:"password"`
//	Age xml.Name `json:"age" xml:"age" yaml:"age" uri:"age"`
//}

func main() {
	r := gin.Default()
	r.POST("someJSON", someJSON)
	r.POST("someStruct", someStruct)
	//r.POST("someXML", someXML)
	r.POST("someYAML", someYAML)
	r.POST("someURI", someURI)
	r.Run(":8080")
}

func someJSON(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"name": user.Name,
		"password": user.Password,
		"age": user.Age,
	})
}

func someStruct(c *gin.Context) {
	var msg struct {
		Name    string
		Message string
		Number  string
	}
	msg.Name = c.Query("name")
	msg.Message = c.Query("message")
	msg.Number = c.Query("number")
	c.JSON(http.StatusOK, msg)
}

//func someXML(c *gin.Context) {
//	var user Message
//	if err := c.ShouldBindXML(&user); err != nil {
//		c.XML(http.StatusBadRequest, gin.H{"error": err.Error()})
//	}
//	c.XML(http.StatusOK,gin.H{
//		"name": user.Name,
//		"password": user.Password,
//		"age": user.Age,
//	})
//}

func someYAML(c *gin.Context) {
	var user User
	if err := c.ShouldBindYAML(&user); err != nil {
		c.YAML(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.YAML(http.StatusOK,gin.H{
		"name": user.Name,
		"password": user.Password,
		"age": user.Age,
	})

}

func someURI(c *gin.Context) {
	var user User
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK,gin.H{
		"name": user.Name,
		"password": user.Password,
		"age": user.Age,
	})

}