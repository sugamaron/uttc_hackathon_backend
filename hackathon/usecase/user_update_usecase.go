package usecase

import (
	"github.com/gin-gonic/gin"
	"hackathon/dao"
	"hackathon/model"
	"log"
	"net/http"
	"strconv"
)

// ユーザ情報更新
func UpdateUser(c *gin.Context) {
	userId := c.Param("user_id")

	//新しいユーザ情報
	var newUser model.User
	newUser.UserId = userId //ユーザidは変更不可とする
	newUser.UserName = c.Query("user_name")
	termInt, _ := strconv.Atoi(c.Query("term"))
	newUser.Term = termInt
	newUser.Email = c.Query("email")

	if err := dao.UpdateUserDao(newUser); err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": newUser})
	}
}
