package controllers

import (
	"net/http"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/courses/repositories"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
	"github.com/gin-gonic/gin"
)

type TestsController struct {
	testsRepository *repositories.TestsRepository
}

func NewTestsController() *TestsController {
	return &TestsController{
		testsRepository: repositories.NewTestsRepository(),
	}
}

func (lc *TestsController) GetTest(ctx *gin.Context) {
	chapterID := ctx.Param("chapter_id")

	authorID := ctx.GetString("author_id")
	userID := ctx.GetString("user_id")

	appendWith := ctx.Query("append_with")

	testsRepository := lc.testsRepository

	if lesson, err := testsRepository.GetTest(chapterID, authorID, userID, appendWith); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, lesson)
	}
}

func (lc *TestsController) CreateQuestion(ctx *gin.Context) {
	testsRepository := lc.testsRepository

	var question repositories.CreatedQuestionDTO

	if err := ctx.ShouldBind(&question); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	chapterID := ctx.GetInt("chapter_id")

	if err := testsRepository.CreateQuestion(uint(chapterID), question); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (lc *TestsController) GetQuestion(ctx *gin.Context) {
	id := ctx.Param("question_id")

	appendWith := ctx.Query("append_with")

	testsRepository := lc.testsRepository

	if lesson, err := testsRepository.GetQuestion(id, appendWith); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, lesson)
	}
}

func (lc *TestsController) UpdateQuestion(ctx *gin.Context) {
	lessonsRepository := lc.testsRepository

	var question repositories.UpdatedQuestionDTO

	if err := ctx.ShouldBind(&question); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	ID := ctx.Param("question_id")

	if err := lessonsRepository.UpdateQuestion(ID, question); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (cc *TestsController) DeleteQuestion(ctx *gin.Context) {
	ID := ctx.Param("question_id")

	testsRepository := cc.testsRepository

	if err := testsRepository.DeleteQuestion(ID); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}