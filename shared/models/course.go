package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Languages string

const (
	StatusPending  Languages = "ar"
	StatusAccepted Languages = "fr"
	StatusRejected Languages = "en"
)

type Level string

const (
	Bigener  Level = "bigener"
	Medium   Level = "medium"
	Advanced Level = "advanced"
)

type Course struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Title       string         `gorm:"unique;not null" json:"title"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Language    Languages      `gorm:"default:'en'" json:"language"`
	Level       Level          `gorm:"default:'bigener'" json:"level"`
	Duration    time.Duration  `gorm:"type:interval" json:"duration"`
	AuthorID    *string        `json:"author_id,omitempty"`
	Author      *Author        `json:"author,omitempty"`
	Categories  []Category     `gorm:"many2many:course_categories;" json:"categories,omitempty"`
	Chapters    []Chapter      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"chapters,omitempty"`
	Learners    []User         `gorm:"many2many:course_learners;association_foreignkey:LearnerID" json:"learners,omitempty"`
}

type Category struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `gorm:"unique;not null" json:"name"`
	Course    []Course       `gorm:"many2many:course_categories;" json:"courses,omitempty"`
}

type CourseCategory struct {
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CourseID   uint           `gorm:"primaryKey" json:"course_id"`
	CategoryID uint           `gorm:"primaryKey" json:"category_id"`
	Course     *Course        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"course"`
	Category   *Category      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"category"`
}

func (CourseCategory) TableName() string {
	return "course_categories"
}

type Chapter struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CourseID    uint           `json:"course_id,omitempty"`
	Course      *Course        `json:"course,omitempty"`
	Lessons     []Lesson       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"lessons,omitempty"`
	Test        *Test          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"test,omitempty"`
}

type Lesson struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	Content     gin.H          `gorm:"json" json:"content,omitempty"`
	Video       *File          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"video,omitempty"`
	ChapterID   uint           `json:"chapter_id,omitempty"`
	Chapter     *Chapter       `json:"chapter,omitempty"`
	Learners    []User         `gorm:"many2many:lesson_learners;association_foreignkey:LearnerID" json:"learners,omitempty"`
}

type LessonLearner struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	LessonID  uint           `gorm:"primaryKey" json:"course_id"`
	LearnerID uint           `gorm:"primaryKey" json:"learner_id"`
	Lesson    *Lesson        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"lesson"`
	Learner   *User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"learner"`
	Learned   bool           `json:"language"`
}

type LearningStatus string

const (
	Succeed  LearningStatus = "succeed"
	Failed   LearningStatus = "failed"
	Learning LearningStatus = "learning"
)

type CourseLearner struct {
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CourseID      uint           `gorm:"primaryKey" json:"course_id"`
	LearnerID     uint           `gorm:"primaryKey" json:"learner_id"`
	Course        *Course        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"course"`
	Learner       *User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"learner"`
	LeaningStatus LearningStatus `gorm:"default:'learning'" json:"language"`
	Rate          *float64       `json:"rate"`
}
