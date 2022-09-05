package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Lange struct {
	Lang string `json:"lang"`
	Tag  string `json:"tag"`
}

// 使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON
func main() {
	r := gin.Default()
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})
	r.Run(":8080")
}

// 转义 json 中的特殊字符
//func main() {
//	r := gin.Default()
//	r.GET("/someJSON", func(c *gin.Context) {
//		data := map[string]interface{}{
//			"Lang": "GO语言",
//			"Tag":  "<br>",
//		}
//		datas, err := json.Marshal(&data)
//		if err != nil {
//			c.JSON(http.StatusBadRequest, err.Error())
//			return
//		}
//		dataa := string(datas)
//		fmt.Println(dataa)
//		bf := bytes.NewBuffer([]byte{})
//		jsonEncoder := json.NewEncoder(bf)
//		jsonEncoder.SetEscapeHTML(false)
//		if err := jsonEncoder.Encode(data); err != nil {
//			c.JSON(http.StatusBadRequest, err.Error())
//			return
//		}
//		fmt.Println(bf.String())
//
//		c.String(http.StatusOK, bf.String())
//	})
//	r.Run(":8080")
//}
