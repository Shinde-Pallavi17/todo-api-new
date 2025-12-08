package models

import "time"

type Task struct {
	Id          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"optional"`
	TaskGroup   string    `json:"task_group" example:"personal" enums:"personal,office,shopping,family,friends,education,health,travel,food" gorm:"type:enum('personal','office','shopping','family','friends','education','health','travel','food')"`
	DueDate     time.Time `json:"due_date" gorm:"not null"`
	Status      string    `json:"status" gorm:"default=pending'"`
	UserID      uint      `json:"user_id" gorm:"foreignKey:UserID"`
	IsOverdue   bool      `json:"is_overdue" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
