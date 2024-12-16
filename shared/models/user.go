package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                string         `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time      `json:"-"`
	UpdatedAt         time.Time      `json:"-"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
	Email             string         `gorm:"unique;not null" json:"email"`
	Password          string         `json:"password"`
	FullName          string         `gorm:"not null" json:"full_name"`
	EmailVerified     bool           `json:"email_verified"`
	Image             *File          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"image,omitempty"`
	AuthorProfile     *Author        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"author_profile,omitempty"`
	Courses          []Course        `gorm:"many2many:course_learners;joinForeignKey:LearnerID;joinReferences:CourseID" json:"courses,omitempty"`
	Lessons           []Lesson       `gorm:"many2many:lesson_learners;" json:"lessons,omitempty"`
	Tests             []Test         `gorm:"many2many:test_results;" json:"tests,omitempty"`
}

type Author struct {
	ID              string         `gorm:"primaryKey" json:"id"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
	Bio             gin.H          `gorm:"json" json:"bio"`
	Balance         float64        `gorm:"balance" json:"balance"`
	UserID          string         `json:"user_id"`
	User            *User          `json:"user"`
	Accomplishments []File         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"accomplishments,omitempty"`
	Courses         []Course       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"courses,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}

func (a *Author) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New().String()
	return nil
}
