package usecase

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hackathon/dao"
	"hackathon/model"
	"io"
	"log"
	"net/http"
)

// ユーザ情報更新
func UpdateUser(c *gin.Context) {
	userId := c.Param("user_id")

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("fail: io.ReadALL, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	var newUser model.User
	if err := json.Unmarshal(body, &newUser); err != nil {
		log.Printf("fail:json.Unmarshal , %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	if err := dao.UpdateUserDao(userId, newUser); err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
