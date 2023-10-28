package usecase

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hackathon/dao"
	"hackathon/model"
	"log"
	"net/http"
)

func GetItems(c *gin.Context) {
	lessonId := c.Query("lesson_id")
	categoryId := c.Query("category_id")
	order := c.Query("order")

	rows, err := dao.GetItemsDao(lessonId, categoryId, order)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	items := make([]model.Item, 0)
	for rows.Next() {
		var i model.Item
		if err := rows.Scan(&i.Title, &i.Registrant, &i.RegistrationDate, &i.UpdateDate, &i.Likes); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		items = append(items, i)
	}

	bytes, err := json.Marshal(items)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", bytes)
}
