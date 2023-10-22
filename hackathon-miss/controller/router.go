package controller

import (
	"github.com/gin-gonic/gin"
	"uttc_hackathon_backend/usecase"
)

func GetRouter() *gin.Engine {
	r = gin.Default()

	r.POST("/users", usecase.RegisterUser)
	r.GET("/users/:user_id", GetUserName)
	r.DELETE("/users/:user_id", DeleteUser)
	r.PUT("/users/:user_id", UpdateUser)
	r.GET("/lessons", GetLessons)
	r.GET("/lessons/:lesson_id", GetLessonName)
	r.GET("/categories", GetCategories)

}
