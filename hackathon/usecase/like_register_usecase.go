package usecase

import "github.com/gin-gonic/gin"

func RegisterLike(c *gin.Context) {
	userId := c.Param("user_id")
	itemId := c.Param("item_id")

}
