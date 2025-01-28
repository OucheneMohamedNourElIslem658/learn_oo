package controllers

import (
	"net/http"
	"strconv"
	"time"  
	"errors"

	"fmt"
	"gorm.io/gorm"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/certaficates/repositories"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	"github.com/gin-gonic/gin"
)

type TestSubmissionRequest struct {
	Answers []struct {
		QuestionID       uint `json:"question_id" binding:"required"`
		SelectedOptionID uint `json:"selected_option_id" binding:"required"`
	} `json:"answers" binding:"required"`
}

type UserProgressController struct {
	userProgressRepo *repositories.UserProgressRepository
}

func NewUserProgressController() *UserProgressController {
	// Pass the database instance here directly
	return &UserProgressController{
		userProgressRepo: repositories.NewUserProgressRepository(database.Instance), 
	}
}

func (c *UserProgressController) CheckCourseCompletion(ctx *gin.Context) {
	courseID := ctx.Param("courseID")
	userID, exists := ctx.Get("id") // Retrieve user ID from the context

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Convert string parameters to uint
	courseIDUint, err := strconv.ParseUint(courseID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	userIDUint := userID.(uint)

	success, apiErr := c.userProgressRepo.HasUserSucceededAllTests(uint(courseIDUint), userIDUint)
	if apiErr != nil {
		ctx.JSON(apiErr.StatusCode, gin.H{"error": apiErr.Message})
		return
	}

	if success {
		// Fetch the user, course, and test results for the certificate
		var user models.User
		var course models.Course
		var testResults []models.TestResult

		if err := c.userProgressRepo.Database.First(&user, userIDUint).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := c.userProgressRepo.Database.First(&course, courseIDUint).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := c.userProgressRepo.Database.Where("learner_id = ? AND has_succeed = true", userIDUint).Find(&testResults).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		certificateData := gin.H{
			"course":      course,
			"user":        user,
			"testResults": testResults,
			"date":        time.Now(),
		}
		ctx.JSON(http.StatusOK, certificateData)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "User has not succeeded all tests"})
	}
}






func (ctrl *UserProgressController) MarkLessonsAsLearned(ctx *gin.Context) {
	userID, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Convert userID to uint
	userIDUint, ok := userID.(uint)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}

	// Extract chapter ID from request parameters and convert it to uint
	chapterIDStr := ctx.Param("chapterID")
	chapterID, err := strconv.ParseUint(chapterIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
		return
	}

	// Call the repository to mark the lessons as learned
	apiErr := ctrl.userProgressRepo.MarkLessonsAsLearned(userIDUint, uint(chapterID))
	if apiErr != nil {
		ctx.JSON(apiErr.StatusCode, gin.H{"error": apiErr.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Lessons marked as learned"})
}


func (upc *UserProgressController) GetTestByChapter(c *gin.Context) {
    // Parse chapter ID
    chapterID, err := strconv.ParseUint(c.Param("chapter_id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
        return
    }

    // Get user ID from context
    userIDInterface, exists := c.Get("id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
        return
    }

    // Convert user ID to uint 
    userID, ok := userIDInterface.(uint)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
        return
    }

    // Convert userID to string (if repository expects string)
    userIDStr := strconv.FormatUint(uint64(userID), 10)

    // Check if chapter exists and get associated course
    chapter, err := upc.userProgressRepo.GetChapterWithCourse(uint(chapterID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
        return
    }

    // Check course access FIRST before fetching test
    hasAccess, err := upc.userProgressRepo.CheckUserCourseAccess(userIDStr, chapter.CourseID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify course access"})
        return
    }
    if !hasAccess {
        c.JSON(http.StatusForbidden, gin.H{"error": "You don't have access to this course"})
        return
    }

    // Get test with questions and options
    test, err := upc.userProgressRepo.GetTestByChapterID(uint(chapterID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Test not found for this chapter"})
        return
    }

    // Handle test attempt
    currentChance, err := upc.userProgressRepo.HandleTestAttempt(userIDStr, test.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to handle test attempt"})
        return
    }

    if currentChance > test.MaxChances {
        c.JSON(http.StatusForbidden, gin.H{"error": "Maximum chances exceeded"})
        return
    }

    // Return response with calculated fields
    c.JSON(http.StatusOK, gin.H{
        "test":           test,
        "current_chance": currentChance,
        "max_chances":    test.MaxChances,
    })
}



func (upc *UserProgressController) SubmitTestAnswers(c *gin.Context) {
	// Get test ID from URL params
	testID, err := strconv.ParseUint(c.Param("test_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test ID"})
		return
	}

	// Get user ID from context
	userIDInterface, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	// Parse request body
	var userAnswers []struct {
		QuestionID       uint `json:"question_id" binding:"required"`
		SelectedOptionID uint `json:"selected_option_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&userAnswers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}

	// Get test with questions and options
	test, err := upc.userProgressRepo.GetTestWithQuestionsAndOptions(uint(testID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Test not found"})
		return
	}

	// Create a map for faster question lookup
	questionMap := make(map[uint]*models.Question)
	for i := range test.Questions {
		questionMap[test.Questions[i].ID] = &test.Questions[i]
	}

	// Create a map for faster option lookup
	optionMap := make(map[uint]map[uint]*models.Option)
	for _, question := range test.Questions {
		optionMap[question.ID] = make(map[uint]*models.Option)
		for i := range question.Options {
			optionMap[question.ID][question.Options[i].ID] = &question.Options[i]
		}
	}

	var correctAnswersCount int
	totalQuestions := len(test.Questions)

	if totalQuestions == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No questions found in test"})
		return
	}

	// Check existing test result
	existingResult, err := upc.userProgressRepo.GetTestResult(userID, test.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve test result"})
		return
	}

	// Compare answers
	for _, answer := range userAnswers {
		question, exists := questionMap[answer.QuestionID]
		if !exists {
			continue
		}

		if optionMap[question.ID] != nil {
			if selectedOption, exists := optionMap[question.ID][answer.SelectedOptionID]; exists {
				if selectedOption.IsCorrect {
					correctAnswersCount++
				}
			}
		}
	}

	// Calculate percentage and determine success
	percentage := float64(correctAnswersCount) / float64(totalQuestions) * 100
	hasSucceed := percentage >= 70

	// Prepare test result
	if existingResult != nil {
		// Update existing result
		existingResult.HasSucceed = hasSucceed
		existingResult.CurrentChance++
		err = upc.userProgressRepo.UpdateTestResult(existingResult)
	} else {
		// Create new result
		newTestResult := &models.TestResult{
			TestID:        test.ID,
			LearnerID:     userID,
			HasSucceed:    hasSucceed,
			CurrentChance: 1,
		}
		err = upc.userProgressRepo.SaveTestResult(newTestResult)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save test result"})
		return
	}

	// Mark lessons as learned if successful
	if hasSucceed {
		if err := upc.userProgressRepo.MarkLessonsAsLearned(userID, test.ChapterID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark lessons as learned"})
			return
		}
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"correct_answers":  correctAnswersCount,
		"total_questions":  totalQuestions,
		"percentage":       percentage,
		"success":          hasSucceed,
		"message":          fmt.Sprintf("You %s the test. Your score: %.2f%%", map[bool]string{true: "passed", false: "did not pass"}[hasSucceed], percentage),
	})
}




func (upc *UserProgressController) GetTestResult(c *gin.Context) {
	// Get user ID from context
	userIDInterface, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	// Get test ID from URL params
	testID, err := strconv.ParseUint(c.Param("test_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test ID"})
		return
	}

	// Retrieve the test result from the repository
	testResult, err := upc.userProgressRepo.GetTestResult(userID, uint(testID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Test result not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve test result"})
		}
		return
	}

	// Return the test result
	c.JSON(http.StatusOK, gin.H{
		"test_result": testResult,
	})
}