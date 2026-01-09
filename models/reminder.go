package models

import "time"

type Reminder struct {
	ID     uint `json:"id" gorm:"primaryKey;autoIncrement"`
	TaskID uint `json:"task_id" gorm:"not null;index"`
	UserID uint `json:"user_id" gorm:"not null;index"`

	ReminderAt time.Time `json:"reminder_at" gorm:"not null"`
	Status     string    `json:"status" gorm:"default:pending"`
	RetryCount int       `json:"retry_count" gorm:"default:0"`

	TriggeredAt *time.Time `json:"triggered_at"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
