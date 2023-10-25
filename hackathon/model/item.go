package model

type Item struct {
	Title        string `json:"title"`
	Registrant   string `json:"registrant"`
	RegisterDate string `json:"register_date"`
	UpdateDate   string `json:"update_date"`
	Likes        int    `json:"likes"`
}

type ItemDetail struct {
	Title        string `json:"title"`
	Registrant   string `json:"registrant"`
	RegisterDate string `json:"register_date"`
	Updater      string `json:"updater"`
	UpdateDate   string `json:"update_date"`
	Description  string `json:"description"`
	Url          string `json:"url"`
	Likes        int    `json:"likes"`
	Price        int    `json:"price"`
}

type ItemForRegistration struct {
	ItemId       string `json:"item_id"`
	Title        string `json:"title"`
	CategoryId   string `json:"category_id"`
	LessonId     string `json:"lesson_id"`
	Registrant   string `json:"registrant"`
	RegisterDate string `json:"register_date"`
	Updater      string `json:"updater"`
	UpdateDate   string `json:"update_date"`
	Description  string `json:"description"`
	Url          string `json:"url"`
	Likes        int    `json:"likes"`
	Price        int    `json:"price"`
}

type LikeNum struct {
	LikeNumStr string `json:"like_num_str"`
}
