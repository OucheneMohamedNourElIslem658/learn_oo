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
	Order       *int                  `form:"order"`
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
		Order: func() int {
			if lesson.Order == nil {
				return 0
			}
			return *lesson.Order
		}(),
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
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Order       *int   `form:"order"`
	Content     gin.H  `json:"content,omitempty" binding:"required"`
}

func (ar *LessonsRepository) CreateLessonWithContent(chapterID uint, lesson CreatedLessonWithContentDTO) (apiError *utils.APIError) {
	lessonToCreate := models.Lesson{
		ChapterID:   chapterID,
		Title:       lesson.Title,
		Description: lesson.Description,
		Order: func() int {
			if lesson.Order == nil {
				return 0
			}
			return *lesson.Order
		}(),
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

type UpdateLessonDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Order       *int   `form:"order"`
	Content     gin.H  `json:"content,omitempty"`
}

func (ar *LessonsRepository) UpdateLesson(id string, lesson UpdateLessonDTO) (apiError *utils.APIError) {
	database := ar.database

	var existingLesson models.Lesson
	err := database.Where("id = ?", id).Preload("Video").First(&existingLesson).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "lesson not found",
			}
		}
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if lesson.Title != "" {
		existingLesson.Title = lesson.Title
	}

	if lesson.Description != "" {
		existingLesson.Description = lesson.Description
	}

	if lesson.Order != nil {
		existingLesson.Order = *lesson.Order
	}

	if (lesson.Content != nil) && (existingLesson.Video == nil) {
		existingLesson.Content = lesson.Content
	} else if (lesson.Content != nil) && (existingLesson.Video != nil) {
		return &utils.APIError{
			StatusCode: http.StatusBadRequest,
			Message: gin.H{
				"Content": "lesson is a video",
			},
		}
	}

	err = database.Save(&existingLesson).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (ar *LessonsRepository) UpdateLessonVideo(ID int, authorID, chapterID string, video multipart.File) (apiError *utils.APIError) {
	database := ar.database
	filestorage := ar.filestorage

	var existingVideo models.File
	err := database.Where("id = ?", ID).First(&existingVideo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if err == nil {
		if err := database.Where("id = ?", existingVideo.ID).Unscoped().Delete(&existingVideo).Error; err != nil {
			return &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
		if existingVideo.ImageKitID != nil {
			if err := filestorage.DeleteFile(*existingVideo.ImageKitID); err != nil {
				return &utils.APIError{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
				}
			}
		}
	}

	path := fmt.Sprintf("/learn_oo/authors/%v/courses/chapters/%v/videos", authorID, chapterID)
	uploadData, err := filestorage.UploadFile(video, path)
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	newVideo := models.File{
		URL:          uploadData.Url,
		ImageKitID:   &uploadData.FileId,
		ThumbnailURL: &uploadData.ThumbnailUrl,
		LessonID:     &ID,
	}
	if err := database.Create(&newVideo).Error; err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (ar *LessonsRepository) GetLesson(ID, authorID, userID, appendWith string) (lesson *models.Lesson, apiError *utils.APIError) {
	database := ar.database

	query := database.Model(&models.Lesson{})

	validExtentions := utils.GetValidExtentions(
		appendWith,
		"video",
		"chapter",
		"learners",
	)

	for _, extention := range validExtentions {
		query.Preload(extention)
	}

	query.Select("lessons.*, courses.is_completed, courses.author_id").
		Joins("JOIN chapters ON chapters.id = lessons.chapter_id").
		Joins("JOIN courses ON courses.id = chapters.course_id").
		Where("lessons.id = ?", ID).
		Where("courses.author_id = ? OR (courses.is_completed = ? AND ? IN (SELECT learner_id FROM course_learners WHERE course_id = courses.id))",
			authorID, true, userID).
		Order("lessons.order, lessons.updated_at")

	var existingLesson models.Lesson
	err := query.Where("id = ?", ID).First(&existingLesson).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "lesson not found",
			}
		}
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return &existingLesson, nil
}

func (ar *LessonsRepository) DeleteLesson(ID string) (apiError *utils.APIError) {
	database := ar.database

	deleteResult := database.Where("id = ?", ID).First(&models.Lesson{})

	err := deleteResult.Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if deleteResult.RowsAffected == 0 {
		return &utils.APIError{
			StatusCode: http.StatusNotFound,
			Message:    "lesson not found",
		}
	}

	return nil
}
