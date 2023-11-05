package usecase

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hackathon/dao"
	"hackathon/model"
	"log"
	"net/http"
	"time"
)

func GetLikedItems(c *gin.Context) {
	userId := c.Query("user_id")

	//アイテム一覧取得
	rows, err := dao.GetLikedItemsDao(userId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	items := make([]model.LikedItem, 0)
	for rows.Next() {
		var i model.LikedItem
		//データベースからdatetime型を受け取るとき、[]uint8に変換されるので、[]uint8型を受け入れるScan用の変数dをつくる
		var d model.RawDateData
		if err := rows.Scan(&i.ItemId, &i.Title, &i.Registrant, &d.RegistrationDate, &d.UpdateDate, &i.Likes, &i.CategoryId); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		//登録日、更新日を型変換　[]uint8→string→time.Time(UTC)→time.Time(jst)
		jst, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			panic(err)
		}
		registrationDateUTC, err := time.Parse("2006-01-02 15:04:05", string(d.RegistrationDate))
		if err != nil {
			log.Printf("fail: time.Parse, %v\n", err)
			return
		}
		i.RegistrationDate = registrationDateUTC.In(jst)

		updateDateUTC, err := time.Parse("2006-01-02 15:04:05", string(d.UpdateDate))
		if err != nil {
			log.Printf("fail: time.Parse, %v\n", err)
			return
		}
		i.UpdateDate = updateDateUTC.In(jst)

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
