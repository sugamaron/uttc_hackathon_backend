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

func GetItemDetail(c *gin.Context) {
	itemId := c.Param("item_id")
	rows, err := dao.GetItemDetailDao(itemId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	items := make([]model.ItemDetail, 0)
	for rows.Next() {
		var i model.ItemDetail
		//データベースからdatetime型を受け取るとき、[]uint8に変換されるので、[]uint8型を受け入れるScan用の変数dをつくる
		var d model.RawDateData
		if err := rows.Scan(&i.Title, &i.Registrant, &d.RegistrationDate, &i.Updater, &d.UpdateDate,
			&i.Description, &i.Url, &i.Likes, &i.Price); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		//[]uint8型をtime.Time型に変換し、iに代入
		i.RegistrationDate, err = time.Parse("2006-01-02 15:04:05", string(d.RegistrationDate))
		i.UpdateDate, err = time.Parse("2006-01-02 15:04:05", string(d.UpdateDate))
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
