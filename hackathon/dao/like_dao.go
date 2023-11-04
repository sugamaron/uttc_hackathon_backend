package dao

import (
	"database/sql"
	"hackathon/model"
	"log"
)

// 誰が何のアイテムにいいねしたかをlikeテーブルに挿入
func InsertLikeDao(like model.Like) error {
	const sql_insert = "INSERT INTO likes(user_id, item_id) VALUE(?, ?)"
	_, err := db.Exec(sql_insert, like.UserId, like.ItemId)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		return err
	} else {
		return nil
	}
}

// いいねを消す
func DeleteLikeDao(userId string, itemId string) error {
	const sql_delete = "DELETE FROM likes WHERE user_id = ? AND item_id = ?"
	_, err := db.Exec(sql_delete, userId, itemId)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		return err
	} else {
		return nil
	}
}

// あるアイテムのいいね数を数える関数
func CountLikeDao(itemId string) (*sql.Rows, error) {
	const sql_count = "SELECT COUNT(*) FROM likes WHERE item_id = ?"
	rows, err := db.Query(sql_count, itemId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}

// ユーザーが特定のアイテムにいいねしているかどうか判定
func GetLikeDao(userId string, itemId string) (*sql.Rows, error) {
	const sql_count = "SELECT COUNT(*) FROM likes WHERE user_id = ? AND item_id = ?"
	rows, err := db.Query(sql_count, userId, itemId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}
