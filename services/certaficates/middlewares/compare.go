// package middleware

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
// 	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/certaficates/repositories"
// 	"github.com/gin-gonic/gin"
// )

// // CompareAndCalculatePercentage is a middleware that calculates the test score and attaches it to the context.
// func CompareAndCalculatePercentage(progressRepository *repositories.ProgressRepository) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Get test from context (set by a previous handler or middleware)
// 		test, exists := c.Get("test")
// 		if !exists {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Test data not found in context"})
// 			c.Abort()
// 			return
// 		}

// 		testDetails := test.(*models.Test)

// 		// Parse request body
// 		var userAnswers []struct {
// 			QuestionID       uint `json:"question_id"`
// 			SelectedOptionID uint `json:"selected_option_id"`
// 		}

// 		if err := c.ShouldBindJSON(&userAnswers); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
// 			c.Abort()
// 			return
// 		}

// 		// Create a map for faster question lookup
// 		questionMap := make(map[uint]*models.Question)
// 		for i := range testDetails.Questions {
// 			questionMap[testDetails.Questions[i].ID] = &testDetails.Questions[i]
// 		}

// 		// Create a map for faster option lookup
// 		optionMap := make(map[uint]map[uint]*models.Option)
// 		for _, question := range testDetails.Questions {
// 			optionMap[question.ID] = make(map[uint]*models.Option)
// 			for i := range question.Options {
// 				optionMap[question.ID][question.Options[i].ID] = &question.Options[i]
// 			}
// 		}

// 		var correctAnswersCount int
// 		totalQuestions := len(testDetails.Questions)

// 		if totalQuestions == 0 {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "No questions found in test"})
// 			c.Abort()
// 			return
// 		}

// 		// Get the learner ID from the context
// 		learnerIDInterface, exists := c.Get("id")
// 		if !exists {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
// 			c.Abort()
// 			return
// 		}

// 		learnerID, ok := learnerIDInterface.(uint)
// 		if !ok {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
// 			c.Abort()
// 			return
// 		}

// 		// Check existing test result
// 		existingResult, err := progressRepository.GetTestResult(learnerID, testDetails.ID)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve test result"})
// 			c.Abort()
// 			return
// 		}

// 		// Compare answers
// 		for _, answer := range userAnswers {
// 			question, exists := questionMap[answer.QuestionID]
// 			if !exists {
// 				continue
// 			}

// 			if optionMap[question.ID] != nil {
// 				if selectedOption, exists := optionMap[question.ID][answer.SelectedOptionID]; exists {
// 					if selectedOption.IsCorrect {
// 						correctAnswersCount++
// 					}
// 				}
// 			}
// 		}

// 		// Calculate percentage and determine success
// 		percentage := float64(correctAnswersCount) / float64(totalQuestions) * 100
// 		hasSucceed := percentage >= 70

// 		// Prepare test result
// 		if existingResult != nil {
// 			// Update existing result
// 			existingResult.HasSucceed = hasSucceed
// 			existingResult.CurrentChance++
// 			err = progressRepository.UpdateTestResult(existingResult)
// 		} else {
// 			// Create new result
// 			newTestResult := &models.TestResult{
// 				TestID:        testDetails.ID,
// 				LearnerID:     learnerID,
// 				HasSucceed:    hasSucceed,
// 				CurrentChance: 1,
// 			}
// 			err = progressRepository.SaveTestResult(newTestResult)
// 		}

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save test result"})
// 			c.Abort()
// 			return
// 		}

// 		// Attach score details to the context
// 		c.Set("score_details", gin.H{
// 			"correct_answers":  correctAnswersCount,
// 			"total_questions":  totalQuestions,
// 			"percentage":       percentage,
// 			"success":          hasSucceed,
// 			"message":          fmt.Sprintf("You %s the test. Your score: %.2f%%", map[bool]string{true: "passed", false: "did not pass"}[hasSucceed], percentage),
// 		})

// 		c.Next()
// 	}
// }