package repositories

import (
	"fmt"
	"mime/multipart"
	"net/http"
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
