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
	ImageKitID   *string        `json:"image_kit_id"`
	UserID       *string        `json:"user_id"`
	User         *User          `json:"user"`
	AuthorID     *string        `json:"author_id"`
	Author       *Author        `json:"author"`
	LessonID     *int           `json:"lesson_id"`
	Lesson       *Lesson        `json:"lesson"`
	VideoCourseID *uint          `gorm:"index" json:"video_course_id"`
    ImageCourseID *uint          `gorm:"index" json:"image_course_id"`
}
