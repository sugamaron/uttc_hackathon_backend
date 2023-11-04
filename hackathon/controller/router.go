package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hackathon/usecase"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	// CORSミドルウェアの設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "https://uttc-hackathon-frontend-jade.vercel.app"} // 許可するオリジンのリスト
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	r.Use(cors.New(config))

	r.POST("/users", usecase.RegisterUser)
	r.GET("/users/:email", usecase.GetUser)
	r.DELETE("/users/:user_id", usecase.DeleteUser)
	r.PUT("/users/:user_id", usecase.UpdateUser)
	r.GET("/lessons", usecase.GetLessons)
	r.GET("/categories", usecase.GetCategories)
	r.GET("/items", usecase.GetItems)
	r.GET("/items/:item_id", usecase.GetItemDetail)
	r.GET("/items/books/:item_id", usecase.GetBookDetail)
	r.GET("/items/likes", usecase.GetLikedItems)
	r.POST("/items", usecase.RegisterItem)
	r.PUT("/items/:item_id", usecase.UpdateItem)
	r.DELETE("/items/:item_id", usecase.DeleteItem)
	r.POST("/likes", usecase.RegisterLike)
	r.DELETE("/likes", usecase.DeleteLike)
	r.GET("/likes", usecase.GetLike)

	return r
}
