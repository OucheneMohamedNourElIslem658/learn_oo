package utils

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
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

func IsImage(file multipart.File) bool {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return false
	}
	contentType := http.DetectContentType(buffer)
	return strings.HasPrefix(contentType, "image/")
}

func IsVideo(file multipart.File) bool {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return false
	}
	contentType := http.DetectContentType(buffer)
	return strings.HasPrefix(contentType, "video/")
}
