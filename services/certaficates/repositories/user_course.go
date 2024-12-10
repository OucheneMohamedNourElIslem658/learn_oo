package repositories

import (
	"net/http"

	gorm "gorm.io/gorm"

	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/payment"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
	"github.com/gin-gonic/gin"
)

type UserCourseRepository struct {
	database *gorm.DB
	payment  *payment.Payment
}

func NewUserCourseRepository() *UserCourseRepository {
	return &UserCourseRepository{
		database: database.Instance,
		payment:  payment.NewPayment(),
	}
}

type CreatedSessionDTO struct {
	PaymentSuccessUrl string `json:"payment_success_url" binding:"required"`
	PaymentFailUrl    string `json:"payment_fail_url" binding:"required"`
}

func (ucr *UserCourseRepository) StartCourse(userID string, courseID uint, session CreatedSessionDTO) (paymentURL *string, apiError *utils.APIError) {
	database := ucr.database

	learner := false
	err := database.Model(&models.CourseLearner{}).
		Select("count(*) > 0").
		Where("learner_id = ? AND course_id = ?", userID, courseID).
		Scan(&learner).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if learner {
		return nil, &utils.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    "user is already a learner",
		}
	}

	var course models.Course
	err = database.Model(&models.Course{}).Where("id = ?", courseID).First(&course).Error
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

	payment := ucr.payment

	if course.Price > 0 {
		checkout, err := payment.MakePayment(session.PaymentSuccessUrl, session.PaymentFailUrl, userID, course)
		if err != nil {
			return nil, &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}

		return &checkout.CheckoutURL, nil
	}

	courseLearner := models.CourseLearner{
		CourseID:  course.ID,
		LearnerID: userID,
	}

	err = database.Create(&courseLearner).Error
	if err != nil {
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil, nil
}

type CheckoutDTO struct {
	Data struct {
		ID       string  `json:"id"`
		Status   string  `json:"status"`
		Metadata []gin.H `json:"metadata"`
	} `json:"data"`
}

func (ucr *UserCourseRepository) PayForCourse(checkout CheckoutDTO) (apiError *utils.APIError) {
	database := ucr.database

	if checkout.Data.Status != "paid" {
		return &utils.APIError{
			StatusCode: http.StatusNotFound,
			Message:    "course not paid",
		}
	}

	metadata := checkout.Data.Metadata[0]

	courseLearner := models.CourseLearner{
		CourseID:   uint(metadata["course_id"].(float64)),
		LearnerID:  metadata["user_id"].(string),
		CheckoutID: &checkout.Data.ID,
	}

	err := database.Create(&courseLearner).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}
