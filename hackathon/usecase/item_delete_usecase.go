package usecase

import (
	"github.com/gin-gonic/gin"
	"hackathon/dao"
	"log"
	"net/http"
)

func DeleteItem(c *gin.Context) {
	itemId := c.Param("item_id")

	if err := dao.DeleteItemDao(itemId); err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": itemId})
	}
}
