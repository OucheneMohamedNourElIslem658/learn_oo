package models

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	URL           string         `json:"url"`
	Height        int            `json:"height"`
	Width         int            `json:"width"`
	ThumbnailURL  *string        `json:"thumbnail_url"`
	ImageKitID    *string        `json:"image_kit_id"`
	UserID        *string        `json:"user_id,omitempty"`
	User          *User          `json:"user,omitempty"`
	AuthorID      *string        `json:"author_id,omitempty"`
	Author        *Author        `json:"author,omitempty"`
	LessonID      *int           `json:"lesson_id,omitempty"`
	Lesson        *Lesson        `json:"lesson,omitempty"`
	VideoCourseID *uint          `gorm:"index" json:"video_course_id,omitempty"`
	ImageCourseID *uint          `gorm:"index" json:"image_course_id,omitempty"`
}
