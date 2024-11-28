package utils

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var ErrorMessages = map[string]string{
	"required":       "required",
	"email":          "invalid",
	"password":       "its lenght must be greater than 5",
	"couse_duration": "must be more than 5 min",
}

func ValidationErrorResponse(err error) gin.H {
	errors := make(gin.H)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, vErr := range validationErrors {
			switch vErr.Tag() {
			case "oneof":
				allowedValues := vErr.Param()
				errors[vErr.Field()] = fmt.Sprintf("The value must be one of the following: %s", allowedValues)
			case "min":
				minValue := vErr.Param()
				errors[vErr.Field()] = fmt.Sprintf("The value must be at least: %s", minValue)
			default:
				errors[vErr.Field()] = ErrorMessages[vErr.Tag()]
			}
		}
		return errors
	} else {
		return gin.H{
			"request": err.Error(),
		}
	}
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return len(password) >= 5
}

func validateDuration(fl validator.FieldLevel) bool {
	duration := fl.Field().String()

	if duration, err := time.ParseDuration(duration); err != nil {
		return false
	} else {
		return duration.Minutes() >= 5
	}
}

func initPasswordValidator() (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); !ok {
		return errors.New("validator initialization failed")
	} else {
		v.RegisterValidation("password", validatePassword)
		v.RegisterValidation("course_duration", validateDuration)
	}
	return nil
}

func InitValidators() {
	if err := initPasswordValidator(); err != nil {
		log.Fatal(err)
	}
}

func IsImage(fileHeader multipart.FileHeader) bool {
	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".webp"}
	for _, ext := range imageExtensions {
		if strings.HasSuffix(fileHeader.Filename, ext) {
			return true
		}
	}

	return false
}

func IsVideo(fileHeader multipart.FileHeader) bool {
	videoExtensions := []string{".mp4", ".mkv", ".avi", ".mov", ".wmv", ".flv", ".webm"}
	for _, ext := range videoExtensions {
		if strings.HasSuffix(fileHeader.Filename, ext) {
			return true
		}
	}
	return false
}
