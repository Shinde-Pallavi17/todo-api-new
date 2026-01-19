package models

import "time"

type Task struct {
	Id          uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string     `json:"title" gorm:"not null"`
	Description string     `json:"description" gorm:"optional"`
	TaskGroup   string     `json:"task_group" example:"personal" enums:"personal,office,shopping,family,friends,education,health,travel,food" gorm:"type:enum('personal','office','shopping','family','friends','education','health','travel','food')"`
	DueDate     time.Time  `json:"due_date" gorm:"optional"`
	ReminderAt  *time.Time `json:"reminder_at" gorm:"default:null"`

	Priority string `json:"priority" gorm:"default=medium"`
	Status   string `json:"status" gorm:"default=pending'"`
	UserID   uint   `json:"user_id" gorm:"foreignKey:UserID"`

	AssignedByID uint   `json:"assigned_by_id"`
	AssignedBy   string `json:"assigned_by" gorm:"type:varchar(100)"`
	AssignedRole string `json:"assigned_role" gorm:"type:varchar(20)"` // admin / user

	IsOverdue bool      `json:"is_overdue" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
