package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func loginHandler (context *gin.Context) {
	if context.Request.Method == "GET" {
		// 调用context.HTML 渲染模板
		// 状态码、模板名、参数( 用于渲染模板中的 {{}}, 这里我们没有使用模板语法, 所以传个 gin.H{} 即可 )
		context.HTML(200, `index.html`, nil)
	} else {
		// 如果不存在的话, 得到的是空字符串, 但是我们也可以设置默认值, 和Query是类似的
		username := context.PostForm("username")
		password := context.PostForm("password")
		// 如果提交多个值, 我们可以使用PostFormArray获取
		context.String(200, "姓名: %v; 密码: %v;", username, password)
		context.HTML(200, `index.html`, nil)
	}
}

func main() {
	router := gin.Default()
	// 这里很关键, 我们的 login.html 是写在当前目录的 templates 目录中的, 所以必须指定模板所在的目录
	// templates/* 表示从templates目录中加载模板文件
	router.LoadHTMLGlob("index.html")
	router.Any("/login", loginHandler)
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}