package controllers

import (
	"net/http"
	"strconv"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/courses/repositories"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
	"github.com/gin-gonic/gin"
)

type LessonsController struct {
	lessonsRepository *repositories.LessonsRepository
}

func NewLessonsController() *LessonsController {
	return &LessonsController{
		lessonsRepository: repositories.NewLessonsRepository(),
	}
}

func (lc *LessonsController) CreateLessonWithContent(ctx *gin.Context) {
	LessonsRepository := lc.lessonsRepository

	if ctx.ContentType() != "multipart/form-data" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "content type must be multipart/form-data",
		})
		return
	}

	var lesson repositories.CreatedLessonWithContentDTO

	if err := ctx.ShouldBind(&lesson); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	chapterIDString := ctx.Param("chapter_id")
	chapterID, _ := strconv.Atoi(chapterIDString)

	if err := LessonsRepository.CreateLessonWithContent(uint(chapterID), lesson); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (lc *LessonsController) CreateLessonWithVideo(ctx *gin.Context) {
	LessonsRepository := lc.lessonsRepository

	if ctx.ContentType() != "multipart/form-data" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "content type must be multipart/form-data",
		})
		return
	}

	var lesson repositories.CreatedLessonWithVideoDTO

	if err := ctx.ShouldBind(&lesson); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	authorID := ctx.GetString("author_id")

	chapterIDString := ctx.Param("chapter_id")
	chapterID, _ := strconv.Atoi(chapterIDString)

	if err := LessonsRepository.CreateLessonWithVideo(uint(chapterID), authorID, lesson); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (lc *LessonsController) GetLesson(ctx *gin.Context) {
	id := ctx.Param("lesson_id")

	authorID := ctx.GetString("author_id")
	userID := ctx.GetString("user_id")

	appendWith := ctx.Query("append_with")

	lessonsRepository := lc.lessonsRepository

	if lesson, err := lessonsRepository.GetLesson(id, authorID, userID, appendWith); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, lesson)
	}
}

func (lc *LessonsController) UpdateLesson(ctx *gin.Context) {
	lessonsRepository := lc.lessonsRepository

	var lesson repositories.UpdateLessonDTO

	if err := ctx.ShouldBind(&lesson); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	ID := ctx.Param("lesson_id")

	if err := lessonsRepository.UpdateLesson(ID, lesson); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusOK)
}

func (lc *LessonsController) UpdateLessonVideo(ctx *gin.Context) {
	idString := ctx.Param("course_id")
	id, _ := strconv.Atoi(idString)

	image, imageHeader, err := ctx.Request.FormFile("video")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if imageHeader == nil || image == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "video not provided",
		})
		return
	}

	if !utils.IsVideo(*imageHeader) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "the file is not an video",
		})
		return
	}

	authorID := ctx.GetString("author_id")
	chapterID := ctx.Param("chapter_id")

	lessonsRepository := lc.lessonsRepository

	apiError := lessonsRepository.UpdateLessonVideo(id, chapterID, authorID, image)
	if apiError != nil {
		ctx.JSON(apiError.StatusCode, gin.H{
			"message": apiError.Message,
		})
		return
	}

	ctx.Status(http.StatusOK)
}

func (cc *LessonsController) DeleteLesson(ctx *gin.Context) {
	ID := ctx.Param("lesson_id")

	lessonsRepository := cc.lessonsRepository

	if err := lessonsRepository.DeleteLesson(ID); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusOK)
}