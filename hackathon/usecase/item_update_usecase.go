package usecase

import (
	"github.com/gin-gonic/gin"
	"hackathon/dao"
	"hackathon/model"
	"log"
	"net/http"
	"time"
)

func UpdateItem(c *gin.Context) {
	itemId := c.Param("item_id")

	var newItem model.ItemForUpdate

	newTitle := c.Query("title")
	newItem.Title = newTitle
	newUpdater := c.Query("updater")
	newItem.Updater = newUpdater
	newUpdateDate := time.Now()
	newItem.UpdateDate = newUpdateDate
	newDescription := c.Query("description")
	newItem.Description = newDescription
	newUrl := c.Query("url")
	newItem.Url = newUrl

	if err := dao.UpdateItemDao(itemId, newItem); err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
