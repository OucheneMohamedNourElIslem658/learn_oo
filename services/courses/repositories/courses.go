package repositories

import (
	"fmt"
	"mime/multipart"
	"time"

	gorm "gorm.io/gorm"

	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	utils "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
)

type CoursesRepository struct {
	database *gorm.DB
}

func NewAuthRepository() *CoursesRepository {
	return &CoursesRepository{
		database: database.Instance,
	}
}

type CreatedCourseDTO struct {
	Title       string                `form:"title" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Price       float64               `form:"price"`
	Language    models.Languages      `form:"language" binding:"required,oneof='ar' 'fr' 'en'"`
	Level       models.Level          `form:"level" binding:"required,oneof='bigener' 'medium' 'advanced'"`
	Duration    time.Duration         `form:"duration"`
	Video       *multipart.FileHeader `form:"video,omitempty" binding:"required"`
	Image       *multipart.FileHeader `form:"image,omitempty" binding:"required"`
}

func (ar *CoursesRepository) CreateCourse(course CreatedCourseDTO) (apiError *utils.APIError) {
	fmt.Println(course.Title)
	fmt.Println(course.Description)
	fmt.Println(course.Language)
	fmt.Println(course.Level)
	fmt.Println(course.Duration)
	fmt.Println(course.Image.Filename)
	fmt.Println(course.Video.Filename)
	return nil
}
