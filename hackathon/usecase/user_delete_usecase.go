package usecase

import (
	"github.com/gin-gonic/gin"
	"hackathon/dao"
	"log"
	"net/http"
)

func DeleteUser(c *gin.Context) {
	user_id := c.Param("user_id")

	if err := dao.DeleteUserDao(user_id); err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": user_id})
	}
}
