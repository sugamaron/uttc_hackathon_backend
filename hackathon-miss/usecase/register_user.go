package usecase

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"io"
	"log"
	"net/http"
	"unicode/utf8"
	"uttc_hackathon_backend/model"
)

func RegisterUser(c *gin.Context) {
	id := ulid.Make()
	idString := id.String()

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("fail: io.ReadALL, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	var user model.User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Printf("fail:json.Unmarshal , %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	user.UserId = idString

	if user.UserName == "" {
		log.Printf("fail:Name is empty, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	if utf8.RuneCountInString(user.UserName) > 50 {
		log.Printf("fail:Length of name is over 50, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

}
