package utils

import (
	"errors"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var ErrorMessages = map[string]string{
	"required": "required",
	"email":    "invalid",
	"password": "its lenght must be greater than 5",
}

func ValidationErrorResponse(err error) gin.H {
	errors := make(gin.H)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, vErr := range validationErrors {
			errors[vErr.Field()] = ErrorMessages[vErr.Tag()]
		}
	}
	return errors
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return len(password) >= 5
}

func initPasswordValidator() (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); !ok {
		return errors.New("validator initialization failed")
	} else {
		v.RegisterValidation("password", validatePassword)
	}
	return nil
}

func InitValidators()  {
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