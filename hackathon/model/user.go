package model

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type InsertId struct {
	Id string `json:"id"`
}
