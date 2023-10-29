package dao

import (
	"database/sql"
	"hackathon/model"
	"log"
)

func GetItemsDao(lessonId string, categoryId string, order string) (*sql.Rows, error) {
	const sql_get = "SELECT item_id, title, registrant, registration_date, update_date, likes FROM item WHERE lesson_id = ? AND category_id = ? ORDER BY ?"

	//取得データの順番は文字列orderできめる
	var sql_order = ""
	switch order {
	case "registration_order":
		sql_order = "registration_date DESC"
	case "update_order":
		sql_order = "update_date DESC"
	case "likes_order":
		sql_order = "likes ASC"
	default:
		sql_order = "registration_date"
	}

	rows, err := db.Query(sql_get, lessonId, categoryId, sql_order)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}

func GetItemDetailDao(itemId string) (*sql.Rows, error) {
	const sql_get = "SELECT title, registrant, registration_date, updater, update_date, description, url, likes, price FROM item WHERE item_id = ?"
	rows, err := db.Query(sql_get, itemId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}

func InsertItemDao(item model.ItemForRegistration) error {
	const sql_insert = "INSERT INTO item(item_id, title, category_id, lesson_id, registrant, registration_date, description, url, likes, price) VALUE(?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err := db.Exec(sql_insert, item.ItemId, item.Title, item.CategoryId, item.LessonId,
		item.Registrant, item.RegistrationDate, item.Description, item.Url, item.Likes, item.Price)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		return err
	} else {
		return nil
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

func UpdateItemDao(itemId string, newItem model.ItemForUpdate) error {
	if newItem.Title != "" {
		_, err := db.Exec("UPDATE item SET title=? WHERE item_id = ?", newItem.Title, itemId)
		if err != nil {
			log.Printf("fail: db.Exec, %v\n", err)
			return err
		}
	}
	if newItem.Description != "" {
		_, err := db.Exec("UPDATE item SET description=? WHERE item_id = ?", newItem.Description, itemId)
		if err != nil {
			log.Printf("fail: db.Exec, %v\n", err)
			return err
		}
	}
	if newItem.Url != "" {
		_, err := db.Exec("UPDATE item SET url=? WHERE item_id = ?", newItem.Url, itemId)
		if err != nil {
			log.Printf("fail: db.Exec, %v\n", err)
			return err
		}
	}
	if newItem.Likes != -1 {
		_, err := db.Exec("UPDATE item SET likes=? WHERE item_id = ?", newItem.Likes, itemId)
		if err != nil {
			log.Printf("fail: db.Exec, %v\n", err)
			return err
		}
	}
	if newItem.Price != -1 {
		_, err := db.Exec("UPDATE item SET price=? WHERE item_id = ?", newItem.Price, itemId)
		if err != nil {
			log.Printf("fail: db.Exec, %v\n", err)
			return err
		}
	}
	const sql_update = "UPDATE item SET updater = ?, update_date = ? WHERE item_id = ?"
	_, err := db.Exec(sql_update, newItem.Updater, newItem.UpdateDate, itemId)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		return err
	}
	return nil
}
