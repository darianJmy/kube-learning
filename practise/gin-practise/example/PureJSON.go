package main

import "github.com/gin-gonic/gin"

// 通常，JSON 使用 unicode 替换特殊 HTML 字符，例如 < 变为 \u003c。如果要按字面对这些字符进行编码，则可以使用PureJSON。

func main() {
	data := map[string]interface{}{
		"html": "<b>hello, world!</b>",
	}
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		// JSON 会使用 unicode 替换特殊 HTML 字符
		c.JSON(200, data)
	})
	r.GET("/purejson", func(c *gin.Context) {
		// PureJSON 就是把内容按照字面意思输出
		c.PureJSON(200, data)
	})
	r.Run(":8080")
}
