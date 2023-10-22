package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"uttc_hackathon_backend/model"
)

var db *sql.DB

func init() {
	// ①-1
	err := godotenv.Load("mysql/.env_mysql")
	if err != nil {
		panic("Error loading .env file")
	}

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	db, err := sql.Open("mysql", connStr)

	// ①-3
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}

func InsertUser(user model.User) error {
	const sql_insert = "INSERT INTO user(user_id, user_name, email, term) VALUE(?, ?, ?)"
	_, err := db.Exec(sql_insert, user.UserId, user.UserName, user.Email, user.Term)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		return err
	} else {
		return nil
	}
}
