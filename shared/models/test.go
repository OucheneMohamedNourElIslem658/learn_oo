package models

import (
	"time"

	"gorm.io/gorm"
)

type Test struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	QuestionsCount *uint           `gorm:"-:migration;->" json:"questions_count,omitempty"`
	Questions      []Question     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"questions,omitempty"`
	MaxChances     uint           `gorm:"default:1" json:"max_chances"`
	ChapterID      uint           `json:"chapter_id"`
	Chapter        *Chapter       `json:"chapter,omitempty"`
	Learners       []User         `gorm:"many2many:test_results;association_foreignkey:LearnerID" json:"learners,omitempty"`
}

type Question struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	Content     string         `json:"content"`
	Description string         `json:"description"`
	Duration    uint           `json:"duration"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	TestID      uint           `json:"test_id"`
	Test        *Test          `json:"test,omitempty"`
	Options     []Option       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"options,omitempty"`
}

type Option struct {
	ID         uint           `gorm:"primaryKey" json:"-"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	Content    string         `json:"content"`
	IsCorrect  bool           `json:"is_correct"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	QuestionID uint           `json:"question_id"`
	Question   *Question      `json:"question,omitempty"`
}

type TestResult struct {
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	TestID        uint           `gorm:"primaryKey" json:"test_id"`
	LearnerID     uint           `gorm:"primaryKey" json:"learner_id"`
	Test          *Test          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"test,omitempty"`
	Learner       *User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"learner,omitempty"`
	HasSucceed    bool           `json:"has_succeed"`
	CurrentChance uint           `json:"current_chance"`
}
