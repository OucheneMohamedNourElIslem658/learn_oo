package controllers

import (
	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/courses/repositories"
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

func (coursesRepository *CoursesController) CreateCourse(ctx *gin.Context) {
}