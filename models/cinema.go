package models

type Cinema struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
}
