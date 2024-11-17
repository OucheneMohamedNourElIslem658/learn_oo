package models

import (
	"time"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

type User struct {
	ID            string         `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Email         string         `gorm:"unique;not null" json:"email"`
	Password      string         `json:"password"`
	FullName      string         `gorm:"not null" json:"full_name"`
	EmailVerified bool           `json:"email_verified"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}
