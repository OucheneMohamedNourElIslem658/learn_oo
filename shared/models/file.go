package models

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	URL          string         `json:"url"`
	ThumbnailURL *string        `json:"thumbnail_url"`
	ImageKitID   *string        `json:"image_kit_id,omitempty"`
	UserID       *string        `json:"user_id,omitempty"`
	User         *User          `json:"user,omitempty"`
	AuthorID     *string        `json:"author_id,omitempty"`
	Author       *Author        `json:"author,omitempty"`
	LessonID     *int           `json:"lesson_id,omitempty"`
	Lesson       *Lesson        `json:"lesson,omitempty"`
	CourseID     *int           `json:"course_id,omitempty"`
	Course       *Course        `json:"course,omitempty"`
}
