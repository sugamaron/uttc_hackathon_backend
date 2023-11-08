package model

import "time"

type Item struct {
	ItemId           string    `json:"item_id"`
	Title            string    `json:"title"`
	Registrant       string    `json:"registrant"`
	RegistrationDate time.Time `json:"registration_date"`
	UpdateDate       time.Time `json:"update_date"`
	Likes            int       `json:"likes"`
	ImageUrl         string    `json:"image_url"`
}

type RawDateData struct {
	RegistrationDate []uint8
	UpdateDate       []uint8
}

type ItemDetail struct {
	Title            string    `json:"title"`
	Registrant       string    `json:"registrant"`
	RegistrationDate time.Time `json:"registration_date"`
	Updater          string    `json:"updater"`
	UpdateDate       time.Time `json:"update_date"`
	Description      string    `json:"description"`
	Url              string    `json:"url"`
	Likes            int       `json:"likes"`
	ImageUrl         string    `json:"image_url"`
}

type BookDetail struct {
	Price int `json:"price"`
}

type ItemForRegistration struct {
	ItemId           string    `json:"item_id"`
	Title            string    `json:"title"`
	CategoryId       string    `json:"category_id"`
	LessonId         string    `json:"lesson_id"`
	Registrant       string    `json:"registrant"`
	RegistrationDate time.Time `json:"registration_date"`
	Updater          string    `json:"updater"`
	UpdateDate       time.Time `json:"update_date"`
	Description      string    `json:"description"`
	Url              string    `json:"url"`
	Likes            int       `json:"likes"`
	Price            int       `json:"price"`
	ImageUrl         string    `json:"image_url"`
}

type ItemForUpdate struct {
	Title       string    `json:"title"`
	Updater     string    `json:"updater"`
	UpdateDate  time.Time `json:"update_date"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	ImageUrl    string    `json:"image_url"`
}

type LikeNum struct {
	LikeNumStr string `json:"like_num_str"`
}

type LikedItem struct {
	ItemId           string    `json:"item_id"`
	Title            string    `json:"title"`
	Registrant       string    `json:"registrant"`
	RegistrationDate time.Time `json:"registration_date"`
	UpdateDate       time.Time `json:"update_date"`
	Likes            int       `json:"likes"`
	CategoryId       string    `json:"category_id"`
}
