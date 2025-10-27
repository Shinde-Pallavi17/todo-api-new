package models

import "time"

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
