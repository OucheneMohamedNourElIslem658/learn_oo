package controllers

import (
	"net/http"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/courses/repositories"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
	"github.com/gin-gonic/gin"
)

type ObjectivesAndRequirementsController struct {
	objectivesAndRequirementsRepository *repositories.ObjectivesAndRequirementsRepository
}

func NewObjectivesAndRequirementsController() *ObjectivesAndRequirementsController {
	return &ObjectivesAndRequirementsController{
		objectivesAndRequirementsRepository: repositories.NewObjectivesAndRequirementsRepository(),
	}
}

func (oarc *ObjectivesAndRequirementsController) CreateObjective(ctx *gin.Context) {
	courseID := ctx.GetUint("course_id")

	var objective repositories.CreatedObjectiveRequirementDTO

	if err := ctx.ShouldBind(&objective); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	chaptersRepository := oarc.objectivesAndRequirementsRepository

	if err := chaptersRepository.CreateObjective(courseID, objective); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (oarc *ObjectivesAndRequirementsController) DeleteObjective(ctx *gin.Context) {
	ID := ctx.Param("objective_id")

	objectivesAndRequirementsRepository := oarc.objectivesAndRequirementsRepository

	if err := objectivesAndRequirementsRepository.DeleteObjective(ID); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (oarc *ObjectivesAndRequirementsController) CreateRequirement(ctx *gin.Context) {
	courseID := ctx.GetUint("course_id")

	var requirement repositories.CreatedObjectiveRequirementDTO

	if err := ctx.ShouldBind(&requirement); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	chaptersRepository := oarc.objectivesAndRequirementsRepository

	if err := chaptersRepository.CreateRequirement(courseID, requirement); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (oarc *ObjectivesAndRequirementsController) DeleteRequirement(ctx *gin.Context) {
	ID := ctx.Param("requirement_id")

	objectivesAndRequirementsRepository := oarc.objectivesAndRequirementsRepository

	if err := objectivesAndRequirementsRepository.DeleteRequirement(ID); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
