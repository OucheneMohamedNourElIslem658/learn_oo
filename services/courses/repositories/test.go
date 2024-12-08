package repositories

import (
	"net/http"
	"time"

	gorm "gorm.io/gorm"

	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	filestorage "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/file_storage"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	utils "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
)

type TestsRepository struct {
	database    *gorm.DB
	filestorage *filestorage.FileStorage
}

func NewTestsRepository() *TestsRepository {
	return &TestsRepository{
		database:    database.Instance,
		filestorage: filestorage.NewFileStorage(),
	}
}

type CreatedTestDTO struct {
	MaxChances uint `json:"max_chances" binding:"omitempty,min=1"`
}

func (cr *TestsRepository) CreateTest(chapterID uint, test CreatedTestDTO) (apiError *utils.APIError) {
	database := cr.database

	maxChances := func() uint {
		if test.MaxChances == 0 {
			return 1
		} else {
			return 0
		}
	}()

	testToCreate := models.Test{
		ChapterID:  chapterID,
		MaxChances: maxChances,
	}

	err := database.Create(&testToCreate).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

type UpdateTestDTO struct {
	MaxChances uint `json:"max_chances" binding:"omitempty,min=1"`
}

func (cr *TestsRepository) UpdateTest(ID string, test UpdateTestDTO) (apiError *utils.APIError) {
	database := cr.database

	var existingTest models.Test
	err := database.Where("id = ?", ID).First(&existingTest).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "test not found",
			}
		}
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if test.MaxChances != 0 {
		existingTest.MaxChances = test.MaxChances
	}

	err = database.Save(&existingTest).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (cr *TestsRepository) GetTest(ID, authorID, userID string, appendWith string) (test *models.Test, apiError *utils.APIError) {
	database := cr.database

	// Build the query with preloads
	query := database.Model(&models.Test{})
	validExtensions := utils.GetValidExtentions(
		appendWith,
		"questions",
		"learners",
		"chapter",
	)

	for _, extension := range validExtensions {
		query.Preload(extension)
	}

	var existingTest models.Test

	// Check if the user is the author of the course
	isAuthor := false
	err := database.Model(&models.Course{}).
		Select("id").
		Where("author_id = ? AND id = (SELECT chapter.course_id FROM tests JOIN chapters ON chapters.id = tests.chapter_id WHERE tests.id = ?)", authorID, ID).
		Scan(&isAuthor).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	// If the user is the author, fetch the test without restrictions
	if isAuthor {
		err := query.Where("id = ?", ID).First(&test).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, &utils.APIError{
					StatusCode: http.StatusNotFound,
					Message:    "test not found",
				}
			}
			return nil, &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
		return &existingTest, nil
	}

	// Check if the user is a learner
	learner := false
	err = database.Model(&models.CourseLearner{}).
		Select("learner_id").
		Where("learner_id = ? AND course_id = (SELECT chapters.course_id FROM tests JOIN chapters ON chapters.id = tests.chapter_id WHERE tests.id = ?)", userID, ID).
		Scan(&learner).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if !learner {
		return nil, &utils.APIError{
			StatusCode: http.StatusForbidden,
			Message:    "user is neither author nor learner",
		}
	}

	// Check if the user has comleted the lessons of the current chapter
	var allLessonsLearned bool
	err = database.Raw(`
		SELECT NOT EXISTS (
			SELECT 1
			FROM lessons l
			LEFT JOIN lesson_learners ll ON ll.lesson_id = l.id AND ll.learner_id = ?
			WHERE l.chapter_id = (SELECT chapter_id FROM tests WHERE id = ?)
			AND (ll.learned = false OR ll.learned IS NULL)
		) AS all_learned
	`, userID, ID).Scan(&allLessonsLearned).Error

	if err != nil {
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if !allLessonsLearned {
		return nil, &utils.APIError{
			StatusCode: http.StatusForbidden,
			Message:    "user has not completed all lessons in the current chapter",
		}
	}

	err = query.Where("id = ?", ID).First(&test).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "test not found",
			}
		}
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return &existingTest, nil
}

func (cr *TestsRepository) DeleteTest(ID string) (apiError *utils.APIError) {
	database := cr.database

	deleteResult := database.Where("id = ?", ID).Unscoped().Delete(models.Test{})

	if deleteResult.RowsAffected == 0 {
		return &utils.APIError{
			StatusCode: http.StatusNotFound,
			Message:    "test not found",
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

type CreatedQuestionDTO struct {
	Content     string        `json:"content" binding:"required"`
	Description string        `json:"description" binding:"required"`
	Duration    time.Duration `form:"duration" binding:"omitempty,min=10000000"`
	Options     []struct {
		Content   string `json:"content" binding:"required"`
		IsCorrect bool   `json:"is_correct" binding:"required"`
	} `json:"options" binding:"required,question_options_list"`
}

func (cr *TestsRepository) CreateQuestion(testID uint, question CreatedQuestionDTO) (apiError *utils.APIError) {
	database := cr.database

	var options []models.Option

	for _, option := range question.Options {
		options = append(options, models.Option{
			Content:   option.Content,
			IsCorrect: option.IsCorrect,
		})
	}

	questionToCreate := models.Question{
		TestID:      testID,
		Content:     question.Content,
		Description: question.Description,
		Duration:    question.Duration,
		Options:     options,
	}

	err := database.Create(&questionToCreate).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (cr *TestsRepository) GetQuestion(ID string, appendWith string) (question *models.Question, apiError *utils.APIError) {
	database := cr.database

	// Build the query with preloads
	query := database.Model(&models.Question{})
	validExtensions := utils.GetValidExtentions(
		appendWith,
		"test",
		"answered_learners",
	)

	for _, extension := range validExtensions {
		query.Preload(extension)
	}
	query.Preload("Options")

	var existingQuestion models.Question

	err := query.Where("id = ?", ID).First(&existingQuestion).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "question not found",
			}
		}
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return &existingQuestion, nil
}

type UpdatedQuestionDTO struct {
	Content     string `json:"content"`
	Description string `json:"description"`
}

func (cr *TestsRepository) UpdateQuestion(ID string, question UpdatedQuestionDTO) (apiError *utils.APIError) {
	database := cr.database

	var existingQuestion models.Question
	err := database.Where("id = ?", ID).First(&existingQuestion).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "question not found",
			}
		}
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if question.Content != "" {
		existingQuestion.Content = question.Content
	}

	if question.Description != "" {
		existingQuestion.Description = question.Description
	}

	err = database.Save(&existingQuestion).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

type CreatedOptionDTO struct {
	Content   string `json:"content" binding:"required"`
	IsCorrect bool   `json:"is_correct" binding:"required"`
}

func (cr *TestsRepository) CreateOption(questionID uint, option CreatedOptionDTO) (apiError *utils.APIError) {
	database := cr.database

	optionToCreate := models.Option{
		QuestionID: questionID,
		Content:    option.Content,
		IsCorrect:  option.IsCorrect,
	}

	err := database.Create(&optionToCreate).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (cr *TestsRepository) DeleteOption(ID string) (apiError *utils.APIError) {
	database := cr.database

	deleteResult := database.Where("id = ?", ID).Unscoped().Delete(models.Option{})

	if deleteResult.RowsAffected == 0 {
		return &utils.APIError{
			StatusCode: http.StatusNotFound,
			Message:    "option not found",
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

func (cr *TestsRepository) DeleteQuestion(ID string) (apiError *utils.APIError) {
	database := cr.database

	deleteResult := database.Where("id = ?", ID).Unscoped().Delete(models.Question{})

	if deleteResult.RowsAffected == 0 {
		return &utils.APIError{
			StatusCode: http.StatusNotFound,
			Message:    "question not found",
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
