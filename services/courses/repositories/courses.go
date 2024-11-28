package repositories

import (
	"fmt"
	"math"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	gorm "gorm.io/gorm"

	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	filestorage "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/file_storage"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/payment"
	utils "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
)

type CoursesRepository struct {
	database    *gorm.DB
	filestorage *filestorage.FileStorage
	payment     *payment.Payment
}

func NewAuthRepository() *CoursesRepository {
	return &CoursesRepository{
		database:    database.Instance,
		filestorage: filestorage.NewFileStorage(),
		payment:     payment.NewPayment(),
	}
}

type CreatedCourseDTO struct {
	Title       string                `form:"title" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Price       float64               `form:"price" binding:"min=50"`
	Language    models.Languages      `form:"language" binding:"required,oneof='ar' 'fr' 'en'"`
	Level       models.Level          `form:"level" binding:"required,oneof='bigener' 'medium' 'advanced'"`
	Duration    time.Duration         `form:"duration"`
	Video       *multipart.FileHeader `form:"video,omitempty" binding:"required"`
	Image       *multipart.FileHeader `form:"image,omitempty" binding:"required"`
}

func (ar *CoursesRepository) CreateCourse(authorID string, course CreatedCourseDTO) (apiError *utils.APIError) {
	// Upload Image And Video:

	filestorage := ar.filestorage

	image, _ := course.Image.Open()
	defer image.Close()

	video, _ := course.Image.Open()
	defer video.Close()

	message := make(map[string]any)

	if image == nil || !utils.IsImage(*course.Image) {
		message["Image"] = "file not an image"
	}

	if video == nil || !utils.IsVideo(*course.Video) {
		message["Video"] = "file not a video"
	}

	if len(message) != 0 {
		return &utils.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    message,
		}
	}

	imageUploadResult, err := filestorage.UploadFile(image, fmt.Sprintf("/learn_oo/authors/%v/courses/images", authorID))
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	videoUploadResult, err := filestorage.UploadFile(video, fmt.Sprintf("/learn_oo/authors/%v/courses/videos", authorID))
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	// Create Payment Product:

	courseToCreate := models.Course{
		AuthorID:    &authorID,
		Title:       course.Title,
		Description: course.Description,
		Price:       course.Price,
		Duration:    course.Duration,
		Language:    course.Language,
		Level:       course.Level,
		Image: &models.File{
			URL:          imageUploadResult.Url,
			ThumbnailURL: &imageUploadResult.ThumbnailUrl,
		},
		Video: &models.File{
			URL:          videoUploadResult.Url,
			ThumbnailURL: &imageUploadResult.ThumbnailUrl,
		},
	}

	if course.Price > 0 {
		payment := ar.payment
		product, err := payment.CreateProduct(courseToCreate)
		if err != nil {
			return &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
		courseToCreate.PaymentPriceID = &product.PriceID
		courseToCreate.PaymentProductID = &product.ID
	}

	// Create Course:
	database := ar.database

	err = database.Create(&courseToCreate).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (ar *CoursesRepository) GetCourse(ID, appendWith string) (course *models.Course, apiError *utils.APIError) {
	database := ar.database

	query := database.Model(&models.Course{})

	validExtentions := utils.GetValidExtentions(
		appendWith,
		"author",
		"image",
		"video",
		"requirements",
		"objectives",
		"categories",
		"chapters",
		"learners",
	)

	for _, extention := range validExtentions {
		query.Preload(extention)
	}

	var existingCourse models.Course
	err := query.Where("id = ?", ID).First(&existingCourse).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "course not found",
			}
		}
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return &existingCourse, nil
}

type CourseSearchDTO struct {
	Title         string           `form:"title"`
	FreeOrPaid    string           `form:"free_or_paid" binding:"omitempty,oneof='free' 'paid'"`
	Language      models.Languages `form:"language" binding:"omitempty,oneof='ar' 'fr' 'en'"`
	Level         models.Level     `form:"level" binding:"omitempty,oneof='bigener' 'medium' 'advanced'"`
	MinDuration   time.Duration    `form:"min_duration" binding:"min=0"`
	MaxDuration   time.Duration    `form:"max_duration" binding:"min=0"`
	PageSize      uint             `form:"page_size,default=10" binding:"min=1"`
	Page          uint             `form:"page,default=1" binding:"min=1"`
	AppendWith    string           `form:"append_with,omitempty"`
	CategoriesIDs string           `form:"categories_ids,omitempty"`
}

func (ar *CoursesRepository) GetCourses(filters CourseSearchDTO) (courses []models.Course, currentPage, count, maxPages *uint, apiError *utils.APIError) {
	database := ar.database

	query := database.Model(&models.Course{})

	title := filters.Title
	language := filters.Language
	level := filters.Level
	minDuration := filters.MinDuration
	maxDuration := filters.MaxDuration
	appendWith := filters.AppendWith
	freePaid := filters.FreeOrPaid
	pageSize := filters.PageSize
	page := filters.Page

	var categoriesIDs []string
	if len(filters.CategoriesIDs) > 0 {
		categoriesIDs = strings.Split(filters.CategoriesIDs, ",")
	}

	validExtentions := utils.GetValidExtentions(
		appendWith,
		"author",
		"image",
		"video",
		"categories",
	)

	for _, extention := range validExtentions {
		query.Preload(extention)
	}

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	if freePaid == "free" {
		query = query.Where("price = ?", 0)
	} else if freePaid == "paid" {
		query = query.Where("price <> ?", 0)
	}

	if language != "" {
		query = query.Where("language = ?", language)
	}

	if level != "" {
		query = query.Where("level = ?", level)
	}

	if minDuration > 0 {
		query = query.Where("duration >= ?", minDuration.Seconds())
	}

	if maxDuration > 0 {
		query = query.Where("duration <= ?", maxDuration.Seconds())
	}

	if len(categoriesIDs) > 0 {
		query = query.Joins("JOIN course_categories ON course_categories.course_id = courses.id").
			Where("course_categories.category_id IN (?)", categoriesIDs)
	}

	query.Select("courses.*, COALESCE(AVG(course_learners.rate), 0) AS rate").
		Joins("LEFT JOIN course_learners ON course_learners.course_id = courses.id").
		Group("courses.id").
		Order("courses.rate DESC, price DESC, created_at DESC, duration DESC")

	var totalRecords int64
	database.Model(&models.User{}).Count(&totalRecords)
	totalPages := uint(math.Ceil(float64(totalRecords) / float64(pageSize)))

	offset := (page - 1) * pageSize
	query.Limit(int(pageSize)).Offset(int(offset))

	var coursesList []models.Course
	err := query.Find(&coursesList).Error
	if err != nil {
		return nil, nil, nil, nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	coursesListLenght := uint(len(coursesList))

	return coursesList, &page, &coursesListLenght, &totalPages, nil
}
