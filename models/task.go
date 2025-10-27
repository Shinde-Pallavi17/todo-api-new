package models

import "time"

type Task struct {
	Id          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"optional"`
	DueDate     time.Time `json:"due_date" gorm:"not null"`
	Status      string    `json:"status" gorm:"default=pending'"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
