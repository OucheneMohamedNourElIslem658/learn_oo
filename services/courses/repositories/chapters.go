package repositories

import (
	"net/http"

	gorm "gorm.io/gorm"

	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	filestorage "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/file_storage"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	utils "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
)

type ChaptersRepository struct {
	database    *gorm.DB
	filestorage *filestorage.FileStorage
}

func NewChaptersRepository() *ChaptersRepository {
	return &ChaptersRepository{
		database:    database.Instance,
		filestorage: filestorage.NewFileStorage(),
	}
}

type CreatedChapterDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (cr *ChaptersRepository) CreateChapter(courseID uint, chapter CreatedChapterDTO) (apiError *utils.APIError) {
	database := cr.database

	chapterToCreate := models.Chapter{
		CourseID:    courseID,
		Title:       chapter.Title,
		Description: chapter.Description,
	}

	err := database.Create(&chapterToCreate).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

type UpdateChapterDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (cr *ChaptersRepository) UpdateChapter(ID string, chapter UpdateChapterDTO) (apiError *utils.APIError) {
	database := cr.database

	var existingChapter models.Chapter
	err := database.Where("id = ?", ID).First(&existingChapter).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "chapter not found",
			}
		}
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if chapter.Title != "" {
		existingChapter.Title = chapter.Title
	}

	if chapter.Description != "" {
		existingChapter.Description = chapter.Description
	}

	err = database.Save(&existingChapter).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (cr *ChaptersRepository) GetChapter(ID, appendWith string) (chapter *models.Chapter, apiError *utils.APIError) {
	database := cr.database

	query := database.Model(&models.Chapter{})

	validExtentions := utils.GetValidExtentions(
		appendWith,
		"course",
		"lessons",
		"test",
	)

	for _, extention := range validExtentions {
		if extention == "Lessons" {
			query.Preload(extention, func(db *gorm.DB) *gorm.DB {
				return db.Select("lessons.id, lessons.title, lessons.description, lessons.chapter_id, CASE WHEN files.id IS NOT NULL THEN TRUE ELSE FALSE END AS is_video").
					Joins("LEFT JOIN files ON lessons.id = files.lesson_id")
			})
		} else {
			query.Preload(extention)
		}
	}

	var existingChapter models.Chapter
	err := query.Where("id = ?", ID).First(&existingChapter).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "chapter not found",
			}
		}
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return &existingChapter, nil
}

func (cr *ChaptersRepository) DeleteChapter(ID string) (apiError *utils.APIError) {
	database := cr.database

	deleteResult := database.Where("id = ?", ID).Unscoped().Delete(models.Chapter{})

	if deleteResult.RowsAffected == 0 {
		return &utils.APIError{
			StatusCode: http.StatusNotFound,
			Message:    "chapter not found",
		}
	}

	err := deleteResult.Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}
