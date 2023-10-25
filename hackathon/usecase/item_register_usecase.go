package usecase

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"hackathon/dao"
	"hackathon/model"
	"io"
	"log"
	"net/http"
	"strconv"
)

func RegisterItem(c *gin.Context) {
	id := ulid.Make()
	idString := id.String()

	// リクエストボディ読み込む
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("fail: io.ReadALL, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	var newItem model.ItemForRegistration
	if err := json.Unmarshal(body, &newItem); err != nil {
		log.Printf("fail:json.Unmarshal , %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	newItem.ItemId = idString

	//newItem.CategoryIdがカテゴリ名になっているのでidに変換する
	categoryRows, err := dao.GetCategoryIdDao(newItem.CategoryId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	if err := categoryRows.Scan(&newItem.CategoryId); err != nil {
		log.Printf("fail: rows.Scan, %v\n", err)

		if err := categoryRows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
			log.Printf("fail: rows.Close(), %v\n", err)
		}
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	//newItem.lessonIdが章の名前になっているのでidに変換する
	lessonRows, err := dao.GetLessonIdDao(newItem.LessonId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	if err := lessonRows.Scan(&newItem.LessonId); err != nil {
		log.Printf("fail: rows.Scan, %v\n", err)

		if err := categoryRows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
			log.Printf("fail: rows.Close(), %v\n", err)
		}
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	//いいね数をnewItem.Likesに格納
	likeRows, err := dao.CountLikeDao(newItem.ItemId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	//dao.CountLikeDaoはsql.Rowsでいいね数を返す。Scanでいいね数を変数lに文字列として格納し、
	//その後、l.LikeNumStrに格納してあるいいね数をint型に変換してnewItem.Likesに格納
	var l model.LikeNum
	if err := likeRows.Scan(&l.LikeNumStr); err != nil {
		log.Printf("fail: rows.Scan, %v\n", err)

		if err := categoryRows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
			log.Printf("fail: rows.Close(), %v\n", err)
		}
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	newItem.Likes, _ = strconv.Atoi(l.LikeNumStr)

	if err := dao.InsertItemDao(newItem); err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": newItem})
	}
}
