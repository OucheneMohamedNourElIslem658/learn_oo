package controllers

import (
	"net/http"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/courses/repositories"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
	"github.com/gin-gonic/gin"
)

type CoursesController struct {
	coursesRepository *repositories.CoursesRepository
}

func NewCoursesController() *CoursesController {
	return &CoursesController{
		coursesRepository: repositories.NewAuthRepository(),
	}
}

func (cc *CoursesController) CreateCourse(ctx *gin.Context) {
	coursesRepository := cc.coursesRepository

	if ctx.ContentType() != "multipart/form-data" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "content type must be multipart/form-data",
		})
		return
	}

	var course repositories.CreatedCourseDTO

	if err := ctx.ShouldBind(&course); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	authorID := ctx.GetString("author_id")

	if err := coursesRepository.CreateCourse(authorID, course); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusCreated)
}
