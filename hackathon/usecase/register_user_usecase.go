package usecase

import "C"
import (
	"db/dao"
	"db/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"io"
	"log"
	"net/http"
)

func HandlerPost(c *gin.Context) {
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

	if user.Name == "" {
		log.Printf("fail:Name is empty, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	if len(user.Name) > 50 {
		log.Printf("fail:Length of name is over 50, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	if user.Age < 20 || user.Age > 80 {
		log.Printf("fail:Age is inappropriate, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	if dao.InsertUser(c, idString, user) != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	} else {
		var newId model.InsertId
		newId.Id = idString
		bytes, err := json.Marshal(newId)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		c.Data(http.StatusOK, "application/json; charset=utf-8", bytes)

	}
}
