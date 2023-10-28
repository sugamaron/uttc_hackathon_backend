package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 章の名前一覧取得
func GetLessonsDao() (*sql.Rows, error) {
	const sql_get = "SELECT lesson_id, lesson_name FROM lesson ORDER BY lesson_pos ASC"
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
