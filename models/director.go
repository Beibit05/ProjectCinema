package models

type Director struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FullName  string `json:"full_name"`
	CreatedAt string `json:"created_at"`
}
