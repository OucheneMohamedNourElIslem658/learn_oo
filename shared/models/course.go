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
	ID               uint           `gorm:"primaryKey" json:"id"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	Title            string         `gorm:"not null" json:"title"`
	Description      string         `json:"description"`
	Price            float64        `json:"price"`
	PaymentPriceID   *string        `gorm:"unique" json:"payment_price_id"`
	PaymentProductID *string        `gorm:"unique" json:"payment_product_id"`
	Language         Languages      `gorm:"default:'en'" json:"language"`
	Level            Level          `gorm:"default:'bigener'" json:"level"`
	Duration         uint           `gorm:"type:bigint" json:"duration"`
	Rate             float64        `gorm:"-:migration;->" json:"rate"`
	RatersCount      uint           `gorm:"-:migration;->" json:"raters_count"`
	IsCompleted      bool           `json:"is_completed"`
	Requirements     []Requirement  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"requirements,omitempty"`
	Objectives       []Objective    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"objectives,omitempty"`
	Video            *File          `gorm:"foreignKey:VideoCourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"video,omitempty"`
	Image            *File          `gorm:"foreignKey:ImageCourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"image,omitempty"`
	AuthorID         string         `json:"author_id"`
	Author           *Author        `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Categories       []Category     `gorm:"many2many:course_categories;" json:"categories,omitempty"`
	Chapters         []Chapter      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"chapters,omitempty"`
	Learners         []User         `gorm:"many2many:course_learners;joinForeignKey:CourseID;joinReferences:LearnerID" json:"learners,omitempty"`
}

type Category struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `gorm:"unique;not null" json:"name"`
	Course    []Course       `gorm:"many2many:course_categories;" json:"courses,omitempty"`
}

type CourseCategory struct {
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	CourseID   uint           `gorm:"primaryKey" json:"course_id"`
	CategoryID uint           `gorm:"primaryKey" json:"category_id"`
	Course     *Course        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"course,omitempty"`
	Category   *Category      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"category,omitempty"`
}

func (CourseCategory) TableName() string {
	return "course_categories"
}

type Objective struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Content   string         `json:"content"`
	CourseID  uint           `json:"course_id"`
	Course    *Course        `json:"course,omitempty"`
}

type Requirement struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Content   string         `json:"content"`
	CourseID  uint           `json:"course_id"`
	Course    *Course        `json:"course,omitempty"`
}

type Chapter struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CourseID    uint           `json:"course_id"`
	Course      *Course        `json:"course,omitempty"`
	Lessons     []Lesson       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"lessons,omitempty"`
	Test        *Test          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"test,omitempty"`
}

type Lesson struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	IsVideo     bool           `gorm:"-:migration;->" json:"is_video"`
	Content     gin.H          `gorm:"json" json:"content"`
	Video       *File          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"video,omitempty"`
	ChapterID   uint           `json:"chapter_id"`
	Chapter     *Chapter       `json:"chapter,omitempty"`
	Learners    []User         `gorm:"many2many:lesson_learners;joinForeignKey:LessonID;joinReferences:LearnerID" json:"learners,omitempty"`
}

type LessonLearner struct {
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	LessonID  uint           `gorm:"primaryKey" json:"course_id"`
	LearnerID uint           `gorm:"primaryKey" json:"learner_id"`
	Lesson    *Lesson        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"lesson,omitempty"`
	Learner   *User          `gorm:"foreignKey:LearnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"learner,omitempty"`
	Learned   bool           `json:"language"`
}

type LearningStatus string

const (
	Succeed  LearningStatus = "succeed"
	Failed   LearningStatus = "failed"
	Learning LearningStatus = "learning"
)

type CourseLearner struct {
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	CourseID      uint           `gorm:"column:course_id;primaryKey" json:"course_id"`
	LearnerID     string         `gorm:"column:learner_id;primaryKey" json:"learner_id"`
	Course        *Course        `gorm:"foreignKey:CourseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"course,omitempty"`
	Learner       *User          `gorm:"foreignKey:LearnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"learner,omitempty"`
	LeaningStatus LearningStatus `gorm:"default:'learning'" json:"language"`
	Rate          *float64       `json:"rate,omitempty"`
	CheckoutID    *string        `json:"checkout_id,omitempty"`
}
