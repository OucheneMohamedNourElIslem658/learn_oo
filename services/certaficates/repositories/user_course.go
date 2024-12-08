package repositories

import (
	"net/http"

	gorm "gorm.io/gorm"

	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/payment"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
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
	PaymentSuccessURL string `json:"payment_success_url" binding:"required"`
	PaymenFailURL     string `json:"payment_fail_url" binding:"required"`
}

func (ucr *UserCourseRepository) StartCourse(userID string, courseID uint, session CreatedSessionDTO) (paymentURL *string, apiError *utils.APIError) {
	database := ucr.database

	var course models.Course
	err := database.Model(&models.Course{}).Where("id = ?", courseID).First(&course).Error
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

	var user models.User
	err = database.Model(&models.User{}).Where("id = ?", user).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "user not found",
			}
		}
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	payment := ucr.payment

	if course.Price > 0 {
		// store customer id in the checkout metadata
		// hash customer id in the id token and add middleware authorizationWithCustomerCheck
		checkout, err := payment.MakePayment(session.PaymenFailURL, session.PaymentSuccessURL, user.PaymentCustomerID, course)
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
	ProductID  string `json:"id"`
	Status     string `json:"status"`
	CustomerID string `json:"customer_id"`
}

func (ucr *UserCourseRepository) PayForCourse(checkout CheckoutDTO) (apiError *utils.APIError) {
	database := ucr.database

	if checkout.Status != "paid" {
		return &utils.APIError{
			StatusCode: http.StatusNotFound,
			Message:    "course not paid",
		}
	}

	var course models.Course
	err := database.Model(&models.Course{}).Where("payment_product_id = ?", checkout.ProductID).First(&course).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "course not found",
			}
		}
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	var user models.User
	err = database.Model(&models.User{}).Where("payment_customer_id = ?", checkout.CustomerID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "user not found",
			}
		}
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	courseLearner := models.CourseLearner{
		CourseID:  course.ID,
		LearnerID: user.ID,
	}

	err = database.Create(&courseLearner).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}
