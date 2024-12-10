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

func (lc *TestsController) CreateTest(ctx *gin.Context) {
	testsRepository := lc.testsRepository

	var test repositories.CreatedTestDTO

	if err := ctx.ShouldBind(&test); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	chapterID := ctx.GetInt("chapter_id")

	if err := testsRepository.CreateTest(uint(chapterID), test); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (lc *TestsController) GetTest(ctx *gin.Context) {
	id := ctx.Param("test_id")

	authorID := ctx.GetString("author_id")
	userID := ctx.GetString("user_id")

	appendWith := ctx.Query("append_with")

	testsRepository := lc.testsRepository

	if lesson, err := testsRepository.GetTest(id, authorID, userID, appendWith); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, lesson)
	}
}

func (lc *TestsController) UpdateTest(ctx *gin.Context) {
	lessonsRepository := lc.testsRepository

	var test repositories.UpdateTestDTO

	if err := ctx.ShouldBind(&test); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	ID := ctx.Param("test_id")

	if err := lessonsRepository.UpdateTest(ID, test); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusOK)
}

func (cc *TestsController) DeleteTest(ctx *gin.Context) {
	ID := ctx.Param("test_id")

	testsRepository := cc.testsRepository

	if err := testsRepository.DeleteTest(ID); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusOK)
}

func (lc *TestsController) CreateQuestion(ctx *gin.Context) {
	testsRepository := lc.testsRepository

	var question repositories.CreatedQuestionDTO

	if err := ctx.ShouldBind(&question); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	testID := ctx.GetInt("test_id")

	if err := testsRepository.CreateQuestion(uint(testID), question); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (lc *TestsController) GetQuestion(ctx *gin.Context) {
	id := ctx.Param("question_id")

	appendWith := ctx.Query("append_with")

	testsRepository := lc.testsRepository

	if lesson, err := testsRepository.GetQuestion(id, appendWith); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	ID := ctx.Param("question_id")

	if err := lessonsRepository.UpdateQuestion(ID, question); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusOK)
}

func (cc *TestsController) DeleteQuestion(ctx *gin.Context) {
	ID := ctx.Param("question_id")

	testsRepository := cc.testsRepository

	if err := testsRepository.DeleteQuestion(ID); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusOK)
}