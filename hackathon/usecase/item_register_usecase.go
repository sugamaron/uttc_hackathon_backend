package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"io"
	"log"
	"net/http"
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

}
