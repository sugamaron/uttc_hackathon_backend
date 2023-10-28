package usecase

import (
	"github.com/gin-gonic/gin"
	"hackathon/dao"
	"hackathon/model"
	"log"
	"net/http"
	"strconv"
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
	newLikes, err := strconv.Atoi(c.Query("like"))
	if err != nil {
		newItem.Likes = -1 //いいね数が更新されない場合、
		// item.daoで更新されなことが判別できるように、newItem.Likesは-1とする。
	} else {
		newItem.Likes = newLikes
	}
	newPrice, err := strconv.Atoi(c.Query("price"))
	if err != nil {
		newItem.Price = -1 //いいね数と同様
	} else {
		newItem.Price = newPrice
	}

	if err := dao.UpdateItemDao(itemId, newItem); err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
