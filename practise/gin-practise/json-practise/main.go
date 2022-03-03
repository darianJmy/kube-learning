package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Name string `json:"name",binding:"test"`
	Password int `json:"password"`
}

type test1 struct {
	Message string `json:"message"`
}

type GinResp struct {
	Code    int         `json:"code"`
	Resp    interface{} `json:"resp,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   error		`json:"error"`
}

var ctx *gin.Context

func main(){
	r := gin.Default()
	r.POST("/login", login)
	r.Run(":8080")
}

func login(c *gin.Context){
	var login1 Login
	if err := c.ShouldBindJSON(&login1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"name": "", "password": "", "err": err.Error()})
	} else {
		if jixingxing, e := test(login1); e != nil {
			ginresp := resourceok(400, jixingxing.Message, e)
			c.JSON(http.StatusBadRequest, ginresp)
			} else {
			ginresp := resourceok(200, jixingxing.Message, e)
			c.JSON(http.StatusOK, ginresp )
		}
	}

}

func test(login1 Login) (*test1, error) {
	if login1.Name == "admin" {
		return nil, fmt.Errorf("admin is feifa")
	}
	message := fmt.Sprintf("%s+%d", login1.Name, login1.Password)
	return &test1{
		Message: message,
	}, nil
}

func resourceok(status int, re interface{}, e error) *GinResp {
	return &GinResp{
		Code: status,
		Resp: re,
		//Error: e,
	}
}