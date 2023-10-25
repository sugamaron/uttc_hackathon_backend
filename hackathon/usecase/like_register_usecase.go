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

// リクエストのデータ量が少ないからパスパラメータで送信
// ファイルの量が少なくて済む
func RegisterLike(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("fail: io.ReadALL, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	var l model.Like
	if err := json.Unmarshal(body, &l); err != nil {
		log.Printf("fail:json.Unmarshal , %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	if err := dao.InsertLikeDao(l); err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "date": l})
	}
}
