package usecase

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hackathon/dao"
	"hackathon/model"
	"log"
	"net/http"
)

func GetLikedItems(c *gin.Context) {
	userId := c.Query("user_id")

	////あるユーザーがいいねしたアイテムのid一覧取得
	//rows1, err := dao.GetLikedItemsIdDao(userId)
	//if err != nil {
	//	log.Printf("fail: db.Query, %v\n", err)
	//	c.String(http.StatusInternalServerError, "Server Error")
	//	return
	//}
	//itemIds := make([]model.LikedItemId, 0)
	//for rows1.Next() {
	//	var i model.LikedItemId
	//
	//	if err := rows1.Scan(&i.ItemId); err != nil {
	//		log.Printf("fail: rows.Scan, %v\n", err)
	//
	//		if err := rows1.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
	//			log.Printf("fail: rows.Close(), %v\n", err)
	//		}
	//		c.String(http.StatusInternalServerError, "Server Error")
	//		return
	//	}
	//	itemIds = append(itemIds, i)
	//}

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
		//[]uin8から文字列型に変換
		i.RegistrationDate = string(d.RegistrationDate)
		i.UpdateDate = string(d.UpdateDate)
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
