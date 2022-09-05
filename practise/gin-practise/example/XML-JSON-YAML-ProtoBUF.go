package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

// 接口返回的渲染

func main() {
	r := gin.Default()
	r.GET("/someJSON", func(c *gin.Context) {
		// JSON 就是返回一个 JSON 格式数据
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"name"`
			Message string `json:"message"`
			Number  int    `json:"number"`
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		// XML 就是返回一个 XML 格式数据
		//<map>
		//    <message>hey</message>
		//    <status>200</status>
		//</map>
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		// YAML 就是返回一个 YAML 格式数据
		//message: hey
		//status: 200
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 会返回一个文件 someProtoBuf
		c.ProtoBuf(http.StatusOK, data)
	})

	r.Run(":8080")
}
