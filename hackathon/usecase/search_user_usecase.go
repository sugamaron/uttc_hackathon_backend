package usecase

import (
	"db/dao"
	"db/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandlerGet(c *gin.Context) {
	name := c.Query("name") // To be filled
	if name == "" {
		log.Println("fail: name is empty")
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	rows, err := dao.SearchUser(name)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	users := make([]model.User, 0)
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		users = append(users, u)
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", bytes)

}
