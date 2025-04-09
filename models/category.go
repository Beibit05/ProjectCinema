package models

type Category struct {
	Id   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
