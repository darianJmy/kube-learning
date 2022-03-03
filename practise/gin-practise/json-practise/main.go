package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

func main(){
	r := gin.Default()
	r.POST("/login", login)
	r.Run(":8080")
}

func login(c *gin.Context){
	var json Login
	c.ShouldBindJSON(&json)
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"name":json.Name, "password":json.Password})
}