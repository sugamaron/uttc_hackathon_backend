package usecase

import (
	"github.com/gin-gonic/gin"
	"hackathon/dao"
	"log"
	"net/http"
	"strings"
)

func UpdateItem(c *gin.Context) {
	itemId := c.Param("item_id")

	newTitle := c.Query("title")
	newUpdater := c.Query("updater")
	newUpdateDate := c.Query("update_date")
	newDescription := c.Query("description")
	newUrl := c.Query("url")
	newPrice := c.Query("price")
	var newItemData = map[string]string{
		"title":       newTitle,
		"updater":     newUpdater,
		"update_date": newUpdateDate,
		"description": newDescription,
		"url":         newUrl,
		"price":       newPrice,
	}

	//mysqlのコマンドで使うための文字列の作成
	updateItemSlice := make([]string, 0)
	for k, v := range newItemData {
		if v != "" {
			updateItemSlice = append(updateItemSlice, k+"="+v)
		}
	}
	updateStr := strings.Join(updateItemSlice, ",")

	if err := dao.UpdateItemDao(itemId, updateStr); err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
