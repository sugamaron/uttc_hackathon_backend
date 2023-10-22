package main

import (
	"db/controller"
	"db/dao"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := controller.GetRouter()
	router.Run(":8000")

	dao.CloseDBWithSysCall()

}
