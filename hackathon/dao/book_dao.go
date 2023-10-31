package dao

import (
	"database/sql"
	"log"
)

func GetBookDetailDao(itemId string) (*sql.Rows, error) {
	const sql_get = "SELECT price FROM book WHERE item_id = ?"
	rows, err := db.Query(sql_get, itemId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}

func InsertItemBookDao(itemId string, price int) error {
	const sql_insert = "INSERT INTO book(item_id, price) VALUES (?, ?)"
	_, err := db.Exec(sql_insert, itemId, price)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		return err
	} else {
		return nil
	}
}
