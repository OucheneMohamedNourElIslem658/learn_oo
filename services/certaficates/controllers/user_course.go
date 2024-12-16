package controllers

import (
	"net/http"
	"strconv"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/certaficates/repositories"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
	"github.com/gin-gonic/gin"
)

type UserCourseController struct {
	userCourseRepository *repositories.UserCourseRepository
}

func NewUserCourseController() *UserCourseController {
	return &UserCourseController{
		userCourseRepository: repositories.NewUserCourseRepository(),
	}
}

func (ucc *UserCourseController) StartCourse(ctx *gin.Context) {
	courseIDString := ctx.Param("course_id")
	courseID, _ := strconv.Atoi(courseIDString)

	userID := ctx.GetString("id")
	authorID := ctx.GetString("author_id")

	var session repositories.CreatedSessionDTO

	if err := ctx.ShouldBind(&session); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	userCourseRepository := ucc.userCourseRepository

	if paymentURL, err := userCourseRepository.StartCourse(userID, authorID, uint(courseID), session); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	} else if paymentURL != nil {
		ctx.JSON(http.StatusAccepted, gin.H{
			"payment_url": paymentURL,
		})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (ucc *UserCourseController) PayForCourse(ctx *gin.Context) {
	var checkout repositories.CheckoutDTO

	if err := ctx.Bind(&checkout); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userCourseRepository := ucc.userCourseRepository

	if err := userCourseRepository.PayForCourse(checkout); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}