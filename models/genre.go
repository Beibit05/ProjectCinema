package models

type Genre struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	//Films []Film `gorm:"foreignKey:GenreID"`
}
