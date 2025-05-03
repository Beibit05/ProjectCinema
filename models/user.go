package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Username  string    `gorm:"column=name" json:"username"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `gorm:"column=password"json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
