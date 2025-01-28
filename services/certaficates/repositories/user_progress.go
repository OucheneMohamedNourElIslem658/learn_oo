package repositories

import (
	"errors"
	"net/http"
	// "strconv"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
	"gorm.io/gorm"
)

type UserProgressRepository struct {
	Database *gorm.DB
}

func NewUserProgressRepository(database *gorm.DB) *UserProgressRepository {
	return &UserProgressRepository{
		Database: database,
	}
}

func (r *UserProgressRepository) HasUserSucceededAllTests(courseID uint, userID string) (bool, *utils.APIError) {
	var count int64
	err := r.Database.Model(&models.TestResult{}).
		Joins("JOIN tests ON test_results.test_id = tests.id").
		Joins("JOIN chapters ON tests.chapter_id = chapters.id").
		Where("test_results.learner_id = ? AND chapters.course_id = ? AND test_results.has_succeed = true", userID, courseID).
		Count(&count).Error
	if err != nil {
		return false, &utils.APIError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	var totalTests int64
	err = r.Database.Model(&models.Test{}).
		Joins("JOIN chapters ON tests.chapter_id = chapters.id").
		Where("chapters.course_id = ?", courseID).
		Count(&totalTests).Error
	if err != nil {
		return false, &utils.APIError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	return count == totalTests, nil
}

func (r *UserProgressRepository) MarkLessonsAsLearned(userID string, chapterID uint) *utils.APIError {
	var test models.Test
	if err := r.Database.Preload("Learners").Where("chapter_id = ?", chapterID).First(&test).Error; err != nil {
		return &utils.APIError{
			Message:    "Failed to retrieve test for the given chapter",
			StatusCode: http.StatusBadRequest,
		}
	}

	var testResult models.TestResult
	if err := r.Database.Where("test_id = ? AND learner_id = ?", test.ID, userID).First(&testResult).Error; err != nil {
		return &utils.APIError{
			Message:    "Failed to retrieve test result for the user",
			StatusCode: http.StatusBadRequest,
		}
	}

	if !testResult.HasSucceed {
		return &utils.APIError{
			Message:    "User did not succeed in the test",
			StatusCode: http.StatusBadRequest,
		}
	}

	var lessons []models.Lesson
	if err := r.Database.Where("chapter_id = ?", chapterID).Find(&lessons).Error; err != nil {
		return &utils.APIError{
			Message:    "Failed to retrieve lessons for the given chapter",
			StatusCode: http.StatusBadRequest,
		}
	}

	for _, lesson := range lessons {
		var lessonLearner models.LessonLearner
		if err := r.Database.Where("lesson_id = ? AND learner_id = ?", lesson.ID, userID).FirstOrCreate(&lessonLearner).Error; err != nil {
			return &utils.APIError{
				Message:    "Failed to find or create LessonLearner record",
				StatusCode: http.StatusBadRequest,
			}
		}

		lessonLearner.Learned = true
		if err := r.Database.Save(&lessonLearner).Error; err != nil {
			return &utils.APIError{
				Message:    "Failed to mark lesson as learned",
				StatusCode: http.StatusBadRequest,
			}
		}
	}

	return nil
}

func (r *UserProgressRepository) GetTestWithQuestionsAndOptions(testID uint) (*models.Test, error) {
	var test models.Test
	err := r.Database.Preload("Questions").
		Preload("Questions.Options").
		First(&test, testID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("test not found")
		}
		return nil, err
	}
	return &test, nil
}

func (r *UserProgressRepository) GetTestResult(learnerID string, testID uint) (*models.TestResult, error) {
	var result models.TestResult
	err := r.Database.Where("learner_id = ? AND test_id = ?", learnerID, testID).
		First(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &result, err
}

func (r *UserProgressRepository) SaveTestResult(result *models.TestResult) error {
	return r.Database.Create(result).Error
}

func (r *UserProgressRepository) UpdateTestResult(result *models.TestResult) error {
	return r.Database.Save(result).Error
}

func (r *UserProgressRepository) GetCurrentAttempts(learnerID string, testID uint) (uint, error) {
	var result models.TestResult
	err := r.Database.Where("learner_id = ? AND test_id = ?", learnerID, testID).
		First(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return result.CurrentChance, nil
}

func (r *UserProgressRepository) GetTestByChapterID(chapterID uint) (*models.Test, error) {
	var test models.Test
	err := r.Database.Preload("Questions").
		Preload("Questions.Options").
		Where("chapter_id = ?", chapterID).
		First(&test).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("test not found")
		}
		return nil, err
	}
	return &test, nil
}

func (r *UserProgressRepository) GetChapterWithCourse(chapterID uint) (*models.Chapter, error) {
	var chapter models.Chapter
	err := r.Database.Preload("Course").First(&chapter, chapterID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("chapter not found")
		}
		return nil, err
	}
	return &chapter, nil
}

func (r *UserProgressRepository) CheckUserCourseAccess(userID string, courseID uint) (bool, error) {
	var courseLearner models.CourseLearner
	err := r.Database.Where("learner_id = ? AND course_id = ?", userID, courseID).First(&courseLearner).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil // User is not a learner in this course
	}
	if err != nil {
		return false, err // Some other error occurred
	}
	return true, nil // User is a learner in the course
}

func (r *UserProgressRepository) HandleTestAttempt(userID string, testID uint) (uint, error) {
	var testResult models.TestResult

	// Check if the test result already exists
	err := r.Database.Where("learner_id = ? AND test_id = ?", userID, testID).First(&testResult).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// No test result found, this is the first attempt, create a new test result
		testResult = models.TestResult{
			LearnerID:     userID,
			TestID:        testID,
			CurrentChance: 1, // First attempt
		}
		err := r.Database.Create(&testResult).Error
		if err != nil {
			return 0, err
		}
		return testResult.CurrentChance, nil
	}

	if err != nil {
		return 0, err
	}

	// If test result exists, increment the current chance
	testResult.CurrentChance++
	err = r.Database.Save(&testResult).Error
	if err != nil {
		return 0, err
	}

	return testResult.CurrentChance, nil
}