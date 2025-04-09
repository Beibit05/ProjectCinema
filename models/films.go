package models

type Film struct {
	Id          uint   `gorm:"primaryKey"  json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Genre       int    `json:"genre"`
	Duration    int    `json:"duration"`
	VideoURL    string `json:"video_url"`
}
