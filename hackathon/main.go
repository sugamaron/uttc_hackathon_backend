package main

import (
	_ "github.com/go-sql-driver/mysql"
	"hackathon/controller"
	"hackathon/dao"
)

func main() {

	router := controller.GetRouter()
	router.Run(":8080")

	dao.CloseDBWithSysCall()

}
