package dao

import (
	"database/sql"
	"log"
)

// カテゴリの名前一覧取得
func GetCategoriesDao() (*sql.Rows, error) {
	const sql_get = "SELECT category_id, category_name FROM category ORDER BY category_pos ASC"
	rows, err := db.Query(sql_get)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}

// カテゴリ名からカテゴリid取得
func GetCategoryIdDao(categoryName string) (*sql.Rows, error) {
	const sql_get = "SELECT category_id FROM category WHERE category_name = ?"
	rows, err := db.Query(sql_get, categoryName)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}
