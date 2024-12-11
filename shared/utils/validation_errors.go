package utils

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	// "github.com/golodash/galidator"
)

var ErrorMessages = map[string]string{
	"required":              "required",
	"email":                 "invalid",
	"password":              "its lenght must be greater than 5",
	"couse_duration":        "must be more than 5 min",
	"price":                 "must be 0 or greater than 50",
	"question_options_list": "must at least contain two elements",
}

func ValidationErrorResponse(err error) gin.H {
	errors := make(gin.H)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, vErr := range validationErrors {
			field := func() string {
				var result []rune
				for i, r := range vErr.Field() {
					if i > 0 && r >= 'A' && r <= 'Z' {
						result = append(result, '_')
					}
					result = append(result, r)
				}
				return strings.ToLower(string(result))
			}()

			switch vErr.Tag() {
			case "oneof":
				allowedValues := vErr.Param()
				errors[field] = fmt.Sprintf("The value must be one of the following: %s", allowedValues)
			case "min":
				minValue := vErr.Param()
				errors[field] = fmt.Sprintf("The value must be at least: %s", minValue)
			default:
				errors[field] = ErrorMessages[vErr.Tag()]
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

func ValidatePrice(fl validator.FieldLevel) bool {
	price := fl.Field().Float()
	return price == 0 || price >= 50
}

func ValidateQuestionOptionsList(fl validator.FieldLevel) bool {
	// if fl.Field().IsNil() || fl.Field().Len() == 0 {
	// 	return true
	// }

	options, ok := fl.Field().Interface().([]struct {
		Content   string `json:"content" binding:"required"`
		IsCorrect *bool   `json:"is_correct" binding:"required"`
	})
	if !ok {
		return false
	}
	fmt.Println("===================")
	fmt.Println(len(options))
	return len(options) >= 2
}

func registerValidators() (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); !ok {
		return errors.New("validator initialization failed")
	} else {
		v.RegisterValidation("password", validatePassword)
		v.RegisterValidation("price", ValidatePrice)
		v.RegisterValidation("question_options_list", ValidateQuestionOptionsList)
	}
	return nil
}

func InitValidators() {
	if err := registerValidators(); err != nil {
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
