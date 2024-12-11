package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                string         `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Email             string         `gorm:"unique;not null" json:"email"`
	Password          string         `json:"password"`
	FullName          string         `gorm:"not null" json:"full_name"`
	EmailVerified     bool           `json:"email_verified"`
	PaymentCustomerID string         `json:"payment_customer_id"`
	Image             *File          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"image"`
	AuthorProfile     *Author        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"author_profile"`
	Courses           []Course       `gorm:"many2many:course_learners;" json:"courses"`
	Lessons           []Lesson       `gorm:"many2many:lesson_learners;" json:"lessons"`
	Tests             []Test         `gorm:"many2many:test_results;" json:"tests"`
}

type Author struct {
	ID              string         `gorm:"primaryKey" json:"id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Bio             gin.H          `gorm:"json" json:"bio"`
	Balance         float64        `gorm:"balance" json:"balance"`
	UserID          string         `json:"user_id"`
	User            *User          `json:"user"`
	Accomplishments []File         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"accomplishments"`
	Courses         []Course       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"courses"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}

func (a *Author) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New().String()
	return nil
}
