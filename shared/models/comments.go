package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Content       string         `json:"content"`
	LessonID      uint           `json:"lesson_id"`
	Lesson        *Lesson        `json:"lesson"`
	UserID        uint           `json:"user_id"`
	User          *User          `json:"user"`
	RepliedTo     *uint          `json:"replied_to"`
	Replies       []Comment      `gorm:"foreignKey:RepliedTo;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"replies"`
	Notifications []Notification `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"notifications"`
}

type Notification struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CommentID   *uint          `json:"comment_id"`
	Comment     *Comment       `json:"comment"`
}
