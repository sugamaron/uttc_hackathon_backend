package usecase

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hackathon/dao"
	"hackathon/model"
	"log"
	"net/http"
)

func GetLike(c *gin.Context) {
	userId := c.Query("user_id")
	itemId := c.Query("item_id")

	rows, err := dao.GetLikeDao(userId, itemId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	var likeNumber int
	for rows.Next() {
		var n model.NumberOfLike
		if err := rows.Scan(&n.LikeNum); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		likeNumber = n.LikeNum
	}

	//ユーザーがアイテムにいいねしていたらtrue、そうでなければfalseを返す
	var result bool
	if likeNumber == 1 {
		result = true
	} else {
		result = false
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", bytes)

}
