package usecase

import (
	"github.com/gin-gonic/gin"
	"hackathon/dao"
	"log"
	"net/http"
)

func DeleteLike(c *gin.Context) {
	userId := c.Param("user_id")
	itemId := c.Param("item_id")

	if err := dao.DeleteLikeDao(userId, itemId); err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "user": userId, "item": itemId})
	}
}
