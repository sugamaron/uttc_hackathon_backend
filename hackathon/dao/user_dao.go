package dao

import (
	"database/sql"
	"fmt"
	"hackathon/model"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

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

func InsertUserDao(user model.User) error {
	const sql_insert = "INSERT INTO user(user_id, user_name, email, term) VALUE(?, ?, ?, ?)"
	_, err := db.Exec(sql_insert, user.UserId, user.UserName, user.Email, user.Term)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		return err
	} else {
		return nil
	}
}

func GetUserDao(user_id string) (*sql.Rows, error) {
	const sql_get = "SELECT user_id, user_name, email, term FROM user WHERE user_id = ?"
	rows, err := db.Query(sql_get, user_id)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	} else {
		return rows, nil
	}
}

func DeleteUserDao(user_id string) error {
	const sql_delete = "DELETE FROM user WHERE user_id = ?"
	_, err := db.Exec(sql_delete, user_id)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		return err
	} else {
		return nil
	}
}

func UpdateUserDao(user model.User) error {
	const sql_update = "UPDATE user SET user_name=?, email=?, term=? WHERE user_id = ?"
	_, err := db.Exec(sql_update, user.UserName, user.Email, user.Term, user.UserId)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		return err
	} else {
		return nil
	}
}

// Ctrl+CでHTTPサーバー停止時にDBをクローズする
func CloseDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
