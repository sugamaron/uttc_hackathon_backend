package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func init() {
	// DB接続のための準備
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
}

func GetCategoriesDao() (*sql.Rows, error) {
	const sql_get = "SELECT category_name FROM category"
	rows, err := db.Query(sql_get)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}

func GetCategoryNameDao(categoryId string) (*sql.Rows, error) {
	const sql_get = "SELECT category_name FROM category WHERE category_id = ?"
	rows, err := db.Query(sql_get, categoryId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}

}
