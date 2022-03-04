package response

import (
	"github.com/gin-gonic/gin"
	"kube-learning/practise/gin-practise/gin-demo/handler"
	"net/http"
)

type GinResp struct {
	Code 	int 		`json:"code"`
	Resp 	interface{} `json:"resp"`
	Message string 		`json:"message"`
	Error 	error    `json:"error"`
}

func (g *GinResp) SetCode(c int) {
	g.Code = c
}

func (g *GinResp) SetResp(r interface{}) {
	g.Resp = r
}

func (g *GinResp) SetMessage(m string) {
	g.Message = m
}

func (g *GinResp) SetError(e error) {
	g.Error = e
}


func GetPractise(c *gin.Context) {
	r := GinResp{}

	r.SetCode(200)

	r.SetMessage("Get Practise Message")

	c.JSON(200, r)
}

func PostPractise(c *gin.Context) {
	r := GinResp{}

	var user handler.User

	if err := c.ShouldBindJSON(&user); err != nil {
		r.SetCode(400)
		r.SetError(err)
		c.JSON(400, r)
		return
	}
	r.SetCode(200)
	r.SetResp(user)
	r.SetMessage("Post Practise Message")
	r.SetError(nil)
	c.JSON(200, r)
}


func CookiePractise(c *gin.Context) {
	cookie := &http.Cookie{
		Name: "abc",
		Value: "123",
		Path: "/",
		MaxAge: 60,
		Secure: false,
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)

	r := GinResp{}
	r.SetCode(200)
	r.SetMessage("Set Cookie")
	r.SetError(nil)
	c.JSON(200, r)
}