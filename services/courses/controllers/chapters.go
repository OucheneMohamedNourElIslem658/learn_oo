package controllers

import (
	"net/http"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/courses/repositories"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
	"github.com/gin-gonic/gin"
)

type ChaptersController struct {
	chaptersRepository *repositories.ChaptersRepository
}

func NewChaptersController() *ChaptersController {
	return &ChaptersController{
		chaptersRepository: repositories.NewChaptersRepository(),
	}
}

func (cc *ChaptersController) CreateChapter(ctx *gin.Context) {
	courseID := ctx.GetUint("course_id")

	var chapter repositories.CreatedChapterDTO

	if err := ctx.ShouldBind(&chapter); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	chaptersRepository := cc.chaptersRepository

	if err := chaptersRepository.CreateChapter(courseID, chapter); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (cc *ChaptersController) GetChapter(ctx *gin.Context) {
	id := ctx.Param("chapter_id")

	appendWith := ctx.Query("append_with")

	chaptersRepository := cc.chaptersRepository

	if chapter, err := chaptersRepository.GetChapter(id, appendWith); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, chapter)
	}
}

func (cc *ChaptersController) UpdateChapter(ctx *gin.Context) {
	var chapter repositories.UpdateChapterDTO

	if err := ctx.ShouldBind(&chapter); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	ID := ctx.Param("chapter_id")

	chaptersRepository := cc.chaptersRepository

	if err := chaptersRepository.UpdateChapter(ID, chapter); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusOK)
}

func (cc *ChaptersController) DeleteChapter(ctx *gin.Context) {
	ID := ctx.Param("chapter_id")

	chaptersRepository := cc.chaptersRepository

	if err := chaptersRepository.DeleteChapter(ID); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusOK)
}
