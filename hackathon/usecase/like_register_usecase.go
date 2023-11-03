package usecase

import (
	"github.com/gin-gonic/gin"
	"hackathon/dao"
	"hackathon/model"
	"log"
	"net/http"
)

func RegisterLike(c *gin.Context) {
	userId := c.Query("user_id")
	itemId := c.Query("item_id")
	log.Println("register_test")

	if err := dao.InsertLikeDao(userId, itemId); err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	//以下、アイテムテーブルのいいね数を更新する
	//アイテムのいいね数を数える
	rows, err := dao.CountLikeDao(itemId)
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
	//アイテムテーブルのいいね数更新
	if err := dao.UpdateLikesDao(itemId, likeNumber); err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "user": userId, "item": itemId})

}
