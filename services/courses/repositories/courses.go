package repositories

import (
	gorm "gorm.io/gorm"

	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	utils "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
)

type CoursesRepository struct {
	database  *gorm.DB
}

func NewAuthRepository() *CoursesRepository {
	return &CoursesRepository{
		database:  database.Instance,
	}
}

func (ar *CoursesRepository) CreateCourse() (apiError *utils.APIError) {
	return nil
}