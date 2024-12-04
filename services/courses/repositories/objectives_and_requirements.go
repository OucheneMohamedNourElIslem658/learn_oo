package repositories

import (
	"net/http"

	gorm "gorm.io/gorm"

	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	utils "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
)

type ObjectivesAndRequirementsRepository struct {
	database *gorm.DB
}

func NewObjectivesAndRequirementsRepository() *ObjectivesAndRequirementsRepository {
	return &ObjectivesAndRequirementsRepository{
		database: database.Instance,
	}
}

type CreatedObjectiveRequirementDTO struct {
	Content string `json:"content" binding:"required"`
}

func (oarr *ObjectivesAndRequirementsRepository) CreateObjective(courseID uint, objective CreatedObjectiveRequirementDTO) (apiError *utils.APIError) {
	database := oarr.database

	objectiveToCreate := models.Objective{
		CourseID: courseID,
		Content:  objective.Content,
	}

	err := database.Create(&objectiveToCreate).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (oarr *ObjectivesAndRequirementsRepository) DeleteObjective(ID string) (apiError *utils.APIError) {
	database := oarr.database

	deleteResult := database.Where("id = ?", ID).Unscoped().Delete(models.Objective{})

	if deleteResult.RowsAffected == 0 {
		return &utils.APIError{
			StatusCode: http.StatusNotFound,
			Message:    "objective not found",
		}
	}

	err := deleteResult.Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (oarr *ObjectivesAndRequirementsRepository) CreateRequirement(courseID uint, requirement CreatedObjectiveRequirementDTO) (apiError *utils.APIError) {
	database := oarr.database

	requirementToCreate := models.Requirement{
		CourseID: courseID,
		Content:  requirement.Content,
	}

	err := database.Create(&requirementToCreate).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (oarr *ObjectivesAndRequirementsRepository) DeleteRequirement(ID string) (apiError *utils.APIError) {
	database := oarr.database

	deleteResult := database.Where("id = ?", ID).Unscoped().Delete(models.Requirement{})

	if deleteResult.RowsAffected == 0 {
		return &utils.APIError{
			StatusCode: http.StatusNotFound,
			Message:    "requirement not found",
		}
	}

	err := deleteResult.Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}