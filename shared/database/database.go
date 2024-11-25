package database

import (
	"fmt"
	"log"
	// "os"
	// "time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
)

var Instance *gorm.DB

func Init() {
	dsn := envs.getDatabaseDSN()

	// sqlLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold: time.Second,
	// 		LogLevel:      logger.Info,
	// 		IgnoreRecordNotFoundError: true,
	// 		Colorful:      true,
	// 	},
	// )

	var err error
	Instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: sqlLogger,
	})
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
		&models.Objective{},
		&models.Requirement{},
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
		&models.QuestionAnswer{},
		&models.QuestionAnswerSelection{},
		&models.File{},
		&models.Comment{},
		&models.Notification{},
	)
	if err != nil {
		return err
	}
	return nil
}
