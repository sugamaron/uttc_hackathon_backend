package controller

import (
	"github.com/gin-gonic/gin"
	"hackathon/usecase"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/users", usecase.RegisterUser)
	r.GET("/users/:user_id", usecase.GetUser)
	r.DELETE("/users/:user_id", usecase.DeleteUser)
	r.PUT("/users/:user_id", usecase.UpdateUser)
	r.GET("/lessons", usecase.GetLessons)
	r.GET("/lessons/:lesson_id", usecase.GetLessonName)
	r.GET("/categories", usecase.GetCategories)
	r.GET("/categories/:category_id", usecase.GetCategoryName)
	r.GET("/items/:lesson_id/:category_id", usecase.GetItems)
	r.GET("/items/:item_id", usecase.GetItemDetail)
	r.POST("/items", usecase.RegisterItem)//未完成
	r.PUT("items/:item_id")//未完成
	r.DELETE("items/:item_id", usecase.DeleteItem)
	r.POST("likes/:user_id/:item_id", usecase.)

	return r
}
