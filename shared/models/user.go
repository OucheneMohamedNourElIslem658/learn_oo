package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Email         string         `gorm:"unique;not null" json:"email"`
	Password      string         `json:"password"`
	FullName      string         `gorm:"not null" json:"full_name"`
	EmailVerified bool          `json:"email_verified"`
}