package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
)

var Instance *gorm.DB

func Init() {
	dsn := envs.getDatabaseDSN()

	var err error
	Instance, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err.Error())
	}

	err = migrateTables()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected succesfully!")
}

func migrateTables() error {
	err := Instance.AutoMigrate(
		&models.User{},
		&models.Author{},
		&models.Course{},
		&models.Category{},
		&models.CourseCategory{},
		&models.Chapter{},
		&models.Lesson{},
		&models.LessonLearner{},
		&models.CourseLearner{},
		&models.Test{},
		&models.Question{},
		&models.Option{},
		&models.TestResult{},
		&models.File{},
		&models.Comment{},
		&models.Notification{},
	)
	if err != nil {
		return err
	}
	return nil
}
