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

	return r
}
