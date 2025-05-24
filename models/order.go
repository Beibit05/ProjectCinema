package models

import "time"

type Order struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id"`
	SessionID  uint      `json:"session_id"` // сеанс ID
	SeatRow    int       `json:"seat_row"`
	SeatNumber int       `json:"seat_number"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}
