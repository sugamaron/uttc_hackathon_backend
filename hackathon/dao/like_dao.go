package dao

import (
	"database/sql"
	"hackathon/model"
	"log"
)

// 誰が何のアイテムにいいねしたかをlikeテーブルに挿入
func InsertLikeDao(l model.Like) error {
	const sql_insert = "INSERT INTO like(user_id, item_id) VALUE(?, ?)"
	_, err := db.Exec(sql_insert, l.UserId, l.ItemId)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		return err
	} else {
		return nil
	}
}

// いいねを消す
func DeleteLikeDao(userId string, itemId string) error {
	const sql_delete = "DELETE FROM like WHERE user_id = ? AND item_id = ?"
	_, err := db.Exec(sql_delete, userId, itemId)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		return err
	} else {
		return nil
	}
}

// いいね数数える関数
func CountLikeDao(itemId string) (*sql.Rows, error) {
	const sql_count = "SELECT COUNT(*) FROM like WHERE item_id = ?"
	rows, err := db.Query(sql_count, itemId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}
