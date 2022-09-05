package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	// 关闭控制台颜色
	gin.DisableConsoleColor()

	// 使用Default() 就是说默认使用（ logger 和 recovery 中间件）创建 gin 路由
	r := gin.Default()

	// 如果使用 New 想添加中间件可以使用 Use
	//r := gin.New()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())

	// 路由组使用 gin.BasicAuth() 中间件
	// 中间件处理就是确认有这个用户后，在请求头里添加
	// Authorization
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets"
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	r.Run(":8080")
}
