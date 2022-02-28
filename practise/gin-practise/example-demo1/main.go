package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.Any("/", WebRoot)
	engine.Run(":8080")
}

func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello,world")
}

