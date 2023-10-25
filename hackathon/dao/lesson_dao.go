package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

// 章の名前一覧取得
func GetLessonsDao() (*sql.Rows, error) {
	const sql_get = "SELECT lesson_name FROM lesson"
	rows, err := db.Query(sql_get)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}

// 特定の章の名前取得
func GetLessonNameDao(lessonId string) (*sql.Rows, error) {
	const sql_get = "SELECT lesson_name FROM lesson WHERE lesson_id = ?"
	rows, err := db.Query(sql_get, lessonId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}

// 章名から章id取得
func GetLessonIdDao(lessonName string) (*sql.Rows, error) {
	const sql_get = "SELECT lesson_id FROM lesson WHERE lesson_name = ?"
	rows, err := db.Query(sql_get, lessonName)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}
