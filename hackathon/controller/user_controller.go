package controller

import (
	"db/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		if c.Request.Method != "GET" && c.Request.Method != "POST" {
			c.String(http.StatusInternalServerError, "Server Error")
			c.Abort() // リクエスト処理を中止
			return
		}
		c.Next() // 次のミドルウェアまたはハンドラを実行
	})

	r.GET("/user", usecase.HandlerGet)
	r.POST("/user", usecase.HandlerPost)

	return r
}
