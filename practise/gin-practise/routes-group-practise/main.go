package main

import (
<<<<<<< HEAD
	"fmt"
	"github.com/gin-gonic/gin"
=======
	"github.com/gin-gonic/gin"
	"io"
>>>>>>> 75e300edf276a3ff29d28f8bd3cc0c10044c423c
	"net/http"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
<<<<<<< HEAD
	v1.GET("/login", login)
	v1.GET("/logout", logout)


=======
	{
		v1.GET("/login", login)
		v1.GET("/logout", logout)
	}
>>>>>>> 75e300edf276a3ff29d28f8bd3cc0c10044c423c
	v2 := r.Group("/v2")
	v2.POST("/postname", postname)
	v2.POST("/postpasswd", postpasswd)

	r.Run(":8080")
}

func login(c *gin.Context) {
<<<<<<< HEAD
	c.String(http.StatusOK, "login")
=======
	c.String(http.StatusOK,"login")
>>>>>>> 75e300edf276a3ff29d28f8bd3cc0c10044c423c
}

func logout(c *gin.Context) {
	c.String(http.StatusOK, "logout")
}

func postname(c *gin.Context) {
<<<<<<< HEAD
	name := c.DefaultPostForm("name", "jixingxing")

	c.String(http.StatusOK, fmt.Sprintf("name:%s", name))

}

func postpasswd(c *gin.Context) {
	passwd := c.DefaultPostForm("passwd", "11111")
	c.String(http.StatusOK, fmt.Sprintf("passwd:%s", passwd))
}

// curl http://127.0.0.1:8080/v2/postname -X POST
=======
	body := c.Request.Body
	x, _ := io.ReadAll(body)
	c.JSON(http.StatusOK, x)
}

func postpasswd(c *gin.Context) {

}
>>>>>>> 75e300edf276a3ff29d28f8bd3cc0c10044c423c
