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
	"time"
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
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	newItem.RegistrationDate = time.Now().In(jst)
	newItem.UpdateDate = time.Now().In(jst)
	newItem.Likes = 0

	//newItem.CategoryIdがカテゴリ名になっているのでidに変換する
	log.Println(newItem.CategoryId)
	categoryRows, err := dao.GetCategoryIdDao(newItem.CategoryId)
	log.Println(categoryRows)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	for categoryRows.Next() {
		if err := categoryRows.Scan(&newItem.CategoryId); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := categoryRows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
	}
	log.Println(newItem.CategoryId)

	//newItem.lessonIdが章の名前になっているのでidに変換する
	lessonRows, err := dao.GetLessonIdDao(newItem.LessonId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	for lessonRows.Next() {
		if err := lessonRows.Scan(&newItem.LessonId); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := categoryRows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
	}

	if err := dao.InsertItemDao(newItem); err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	//カテゴリが技術書の場合のみ、bookテーブルに価格の情報をいれる
	if newItem.CategoryId == "book" {
		if err := dao.InsertItemBookDao(newItem.ItemId, newItem.Price); err != nil {
			log.Printf("fail: db.Exec, %v\n", err)
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": newItem})
}
