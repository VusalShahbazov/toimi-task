package entity

import "time"

type Announcement struct {
	Id          int                 `json:"id" pg:"id,pk"`
	Title       string              `json:"title" pg:"title"`
	Description string              `json:"description" pg:"description"`
	Price       float64             `json:"price" pg:"price"`
	CreatedAt   time.Time           `json:"created_at" pg:"created_at"`
	Photos      []AnnouncementPhoto `json:"photos" pg:"rel:has-many"`
}

type AnnouncementPhoto struct {
	Id             int    `json:"id" pg:"id,pk"`
	AnnouncementId int    `json:"announcement_id" pg:"announcement_id"`
	Url            string `json:"url" pg:"url"`
}
