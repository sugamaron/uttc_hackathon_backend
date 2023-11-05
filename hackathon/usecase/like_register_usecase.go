package usecase

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hackathon/dao"
	"hackathon/model"
	"io"
	"log"
	"net/http"
)

func RegisterLike(c *gin.Context) {

	// リクエストボディ読み込む
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("fail: io.ReadALL, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	var like model.Like
	if err := json.Unmarshal(body, &like); err != nil {
		log.Printf("fail:json.Unmarshal , %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	//いいねテーブルに誰が何のアイテムにいいねしたか保存
	if err := dao.InsertLikeDao(like); err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	//以下、アイテムテーブルのいいね数を更新する
	//アイテムのいいね数を数える
	rows, err := dao.CountLikeDao(like.ItemId)
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
	if err := dao.UpdateLikesDao(like.ItemId, likeNumber); err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": like})

}
