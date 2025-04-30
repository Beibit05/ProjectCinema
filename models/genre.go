package models

import "time"

type Genre struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"unique" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
