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

	//カテゴリまたは章が選択されていないとき、エラーを返す
	if newItem.CategoryId == "notSelected" {
		log.Printf("fail: Category is not selected %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	if newItem.LessonId == "notSelected" {
		log.Printf("fail: Lesson is not selected %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	// 画像urlの欄に何も入力されていない時はno imageの画像パス
	if newItem.ImageUrl == "" {
		newItem.ImageUrl = "http://design-ec.com/d/e_others_50/l_e_others_500.jpg"
	}

	newItem.ItemId = idString
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	newItem.RegistrationDate = time.Now().In(jst)
	newItem.UpdateDate = time.Now().In(jst)
	newItem.Likes = 0

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
