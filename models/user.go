package models

import "time"

type User struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username      string    `gorm:"not null" json:"username"`
	Email         string    `gorm:"unique;not null" json:"email"`
	EmailVerified bool      `json:"email_verified" gorm:"default:true"`
	Password      string    `gorm:"not null" json:"-" `
	Role          string    `json:"role" gorm:"type:varchar(20);default:user"` // user | admin
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	Tasks []Task `gorm:"foreignKey:UserID" json:"tasks,omitempty"`
}
