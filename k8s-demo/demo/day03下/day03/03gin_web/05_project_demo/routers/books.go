package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BookHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Book Router",
		"data": nil,
	})
}
//在main.go初始化gin对象之后，调用此方法，传入gin对象，就注册这个路由
func LoadBooks(r *gin.Engine) {
	r.GET("/book", BookHandler)
}