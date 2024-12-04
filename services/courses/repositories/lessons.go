package repositories

import (
	"fmt"
	"mime/multipart"
	"net/http"

	gorm "gorm.io/gorm"

	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	filestorage "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/file_storage"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/payment"
	utils "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
	"github.com/gin-gonic/gin"
)

type LessonsRepository struct {
	database    *gorm.DB
	filestorage *filestorage.FileStorage
	payment     *payment.Payment
}

func NewLessonsRepository() *LessonsRepository {
	return &LessonsRepository{
		database:    database.Instance,
		filestorage: filestorage.NewFileStorage(),
		payment:     payment.NewPayment(),
	}
}

type CreatedLessonWithVideoDTO struct {
	Title       string                `form:"title" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Video       *multipart.FileHeader `form:"video,omitempty" binding:"required"`
}

func (ar *LessonsRepository) CreateLessonWithVideo(chapterID uint, authorID string, lesson CreatedLessonWithVideoDTO) (apiError *utils.APIError) {
	filestorage := ar.filestorage

	video, _ := lesson.Video.Open()
	defer video.Close()

	message := make(map[string]any)

	if video == nil || !utils.IsVideo(*lesson.Video) {
		message["Video"] = "file not a video"
	}

	if len(message) != 0 {
		return &utils.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    message,
		}
	}

	videoUploadResult, err := filestorage.UploadFile(video, fmt.Sprintf("/learn_oo/authors/%v/courses/chapters/%v/videos", authorID, chapterID))
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	lessonToCreate := models.Lesson{
		ChapterID:   chapterID,
		Title:       lesson.Title,
		Description: lesson.Description,
		Video: &models.File{
			URL:          videoUploadResult.Url,
			ThumbnailURL: &videoUploadResult.ThumbnailUrl,
		},
	}

	database := ar.database
	err = database.Create(&lessonToCreate).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

type CreatedLessonWithContentDTO struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
	Content     gin.H  `form:"content,omitempty" binding:"required"`
}

func (ar *LessonsRepository) CreateLessonWithContent(chapterID uint, lesson CreatedLessonWithContentDTO) (apiError *utils.APIError) {
	lessonToCreate := models.Lesson{
		ChapterID:   chapterID,
		Title:       lesson.Title,
		Description: lesson.Description,
		Content: lesson.Content,
	}

	database := ar.database
	err := database.Create(&lessonToCreate).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

type UpdateLessonWithContentDTO struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
	Content     gin.H  `form:"content,omitempty" binding:"required"`
}