package dao

import (
	"database/sql"
	"hackathon/model"
	"log"
)

func GetItemsDao(lessonId string, categoryId string, order string) (*sql.Rows, error) {
	const sql_get = "SELECT item_id, title, registrant, registration_date, update_date, likes, image_url FROM item WHERE lesson_id = ? AND category_id = ? ORDER BY"

	//取得データの順番は文字列orderできめる
	var sql_order string
	switch order {
	case "registration":
		sql_order = "registration_date DESC"
	case "update":
		sql_order = "update_date DESC"
	case "likes":
		sql_order = "likes DESC"
	default:
		sql_order = "registration_date DESC"
	}
	query := sql_get + " " + sql_order

	rows, err := db.Query(query, lessonId, categoryId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}

func GetItemDetailDao(itemId string) (*sql.Rows, error) {
	const sql_get = "SELECT title, registrant, registration_date, updater, update_date, description, url, likes, image_url FROM item WHERE item_id = ?"
	rows, err := db.Query(sql_get, itemId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}

func InsertItemDao(item model.ItemForRegistration) error {
	const sql_insert = "INSERT INTO item(item_id, title, category_id, lesson_id, registrant, registration_date, updater, update_date, description, url, likes, image_url) VALUE(?,?,?,?,?,?,?,?,?,?,?, ?)"
	_, err := db.Exec(sql_insert, item.ItemId, item.Title, item.CategoryId, item.LessonId,
		item.Registrant, item.RegistrationDate, item.Updater, item.UpdateDate, item.Description, item.Url, item.Likes, item.ImageUrl)
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
	const sql_update = "UPDATE item SET updater = ?, update_date = ? WHERE item_id = ?"
	_, err := db.Exec(sql_update, newItem.Updater, newItem.UpdateDate, itemId)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		return err
	}
	return nil
}

func UpdateLikesDao(itemId string, likeNum int) error {
	const sql_update = "UPDATE item SET likes = ? WHERE item_id = ?"
	_, err := db.Exec(sql_update, likeNum, itemId)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		return err
	}
	return nil
}

func GetLikedItemsDao(userId string) (*sql.Rows, error) {
	const sql_get = "SELECT item.item_id, title, registrant, registration_date, update_date, likes, category_id FROM item INNER JOIN likes ON item.item_id = likes.item_id WHERE likes.user_id = ?"
	rows, err := db.Query(sql_get, userId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}
