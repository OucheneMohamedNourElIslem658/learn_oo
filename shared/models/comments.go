package models

import (
	"time"
)

type Comment struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time      `json:"-"`
	Content       string         `json:"content"`
	LessonID      uint           `json:"lesson_id"`
	Lesson        *Lesson        `json:"lesson,omitempty"`
	UserID        uint           `json:"user_id"`
	User          *User          `json:"user,omitempty"`
	RepliedTo     *uint          `json:"replied_to"`
	Replies       []Comment      `gorm:"foreignKey:RepliedTo;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"replies,omitempty"`
	Notifications []Notification `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"notifications,omitempty"`
}

type Notification struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"-"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CommentID   *uint          `json:"comment_id"`
	Comment     *Comment       `json:"comment,omitempty"`
}
