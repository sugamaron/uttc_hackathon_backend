package usecase

import (
	"github.com/gin-gonic/gin"
	"hackathon/dao"
	"log"
	"net/http"
	"strings"
)

// ユーザ情報更新
func UpdateUser(c *gin.Context) {
	userId := c.Param("user_id")

	newUserName := c.Query("user_name")
	newTerm := c.Query("term")
	newEmail := c.Query("email")
	var newUserData = map[string]string{
		"user_name": newUserName,
		"term":      newTerm,
		"email":     newEmail,
	}

	//mysqlのコマンドで使うための文字列の作成
	updateStrSlice := make([]string, 0)
	for k, v := range newUserData {
		if v != "" {
			updateStrSlice = append(updateStrSlice, k+"="+v)
		}
	}
	updateStr := strings.Join(updateStrSlice, ",")

	if err := dao.UpdateUserDao(userId, updateStr); err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
