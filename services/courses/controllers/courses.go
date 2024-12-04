package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/courses/repositories"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
	"github.com/gin-gonic/gin"
)

type CoursesController struct {
	coursesRepository *repositories.CoursesRepository
}

func NewCoursesController() *CoursesController {
	return &CoursesController{
		coursesRepository: repositories.NewCoursesRepository(),
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

func (cc *CoursesController) GetCourse(ctx *gin.Context) {
	id := ctx.Param("course_id")

	authorID := ctx.GetString("author_id")

	appendWith := ctx.Query("append_with")

	coursesRepository := cc.coursesRepository

	fmt.Println(authorID)

	if course, err := coursesRepository.GetCourse(id, authorID, appendWith); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, course)
	}
}

func (cc *CoursesController) GetCourses(ctx *gin.Context) {
	var filters repositories.CourseSearchDTO

	if err := ctx.ShouldBindQuery(&filters); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": utils.ValidationErrorResponse(err),
		})
		return
	}

	coursesRepository := cc.coursesRepository

	if courses, currentPage, count, maxPages, err := coursesRepository.GetCourses(filters); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"courses":      courses,
			"current_page": currentPage,
			"count":        count,
			"max_pages":    maxPages,
		})
	}
}

func (cc *CoursesController) UpdateCourse(ctx *gin.Context) {
	coursesRepository := cc.coursesRepository

	var course repositories.UpdateCourseDTO

	if err := ctx.ShouldBind(&course); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	ID := ctx.Param("course_id")
	authorID := ctx.GetString("author_id")

	if err := coursesRepository.UpdateCourse(ID, authorID, course); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusOK)
}

func (cc *CoursesController) UpdateCourseImage(ctx *gin.Context) {
	idString := ctx.Param("course_id")
	id, _ := strconv.Atoi(idString)

	image, imageHeader, err := ctx.Request.FormFile("image")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if imageHeader == nil || image == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "image not provided",
		})
		return
	}

	if !utils.IsImage(*imageHeader) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "the file is not an image",
		})
		return
	}

	authorID := ctx.GetString("author_id")

	profilesRepository := cc.coursesRepository

	apiError := profilesRepository.UpdateCourseImage(uint(id), authorID, image)
	if apiError != nil {
		ctx.JSON(apiError.StatusCode, gin.H{
			"message": apiError.Message,
		})
		return
	}

	ctx.Status(http.StatusOK)
}

func (cc *CoursesController) UpdateCourseVideo(ctx *gin.Context) {
	idString := ctx.Param("course_id")
	id, _ := strconv.Atoi(idString)

	video, videoHeader, err := ctx.Request.FormFile("video")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if videoHeader == nil || video == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "image not provided",
		})
		return
	}

	if !utils.IsVideo(*videoHeader) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "the file is not a video",
		})
		return
	}

	authorID := ctx.GetString("author_id")

	profilesRepository := cc.coursesRepository

	apiError := profilesRepository.UpdateCourseVideo(uint(id), authorID, video)
	if apiError != nil {
		ctx.JSON(apiError.StatusCode, gin.H{
			"message": apiError.Message,
		})
		return
	}

	ctx.Status(http.StatusOK)
}

func (cc *CoursesController) DeleteCourse(ctx *gin.Context) {
	ID := ctx.Param("course_id")
	authorID := ctx.GetString("author_id")

	coursesRepository := cc.coursesRepository

	if err := coursesRepository.DeleteCourse(ID, authorID); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusOK)
}

func (cc *CoursesController) GetCategories(ctx *gin.Context) {
	coursesRepository := cc.coursesRepository

	categories, err := coursesRepository.GetCategories()

	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"categories": categories,
		"count":      len(categories),
	})
}

func (cc *CoursesController) CreateCategory(ctx *gin.Context) {
	coursesRepository := cc.coursesRepository

	var category repositories.CreatedCategoryDTO

	if err := ctx.ShouldBind(&category); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	err := coursesRepository.CreateCategory(category)

	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (cc *CoursesController) DeleteCategory(ctx *gin.Context) {
	ID := ctx.Param("category_id")

	coursesRepository := cc.coursesRepository

	if err := coursesRepository.DeleteCategory(ID); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusOK)
}
