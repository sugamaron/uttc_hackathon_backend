package dao

import (
	"database/sql"
	"log"
)

func GetItemsDao(lessonId string, categoryId string) (*sql.Rows, error) {
	const sql_get = "SELECT title, registrant, register_date, update_date, likes FROM item WHERE lesson_id = ? AND category_id = ?"
	rows, err := db.Query(sql_get, lessonId, categoryId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}

func GetItemDetailDao(itemId string) (*sql.Rows, error) {
	const sql_get = "SELECT title, registrant, register_date, updater, update_date, description, url, likes, price FROM item WHERE item_id = ?"
	rows, err := db.Query(sql_get, itemId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}

func DeleteItemDao(itemId string) error {
	const sql_delete = "DELETE FROM item WHERE item_id = ?"
	_, err := db.Exec(sql_delete, itemId)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		return err
	} else {
		return nil
	}
}
