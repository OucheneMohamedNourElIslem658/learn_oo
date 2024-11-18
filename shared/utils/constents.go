package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Object gin.H

func GetValidExtentions(initialValues string, validValues ...string) []string {
	extentions := strings.Split(initialValues, ",")
	validExtentions := make([]string, 0)
	for _, extention := range extentions {
		extention = strings.ToLower(extention)
		isExtentionValid := Contains(validValues, extention)
		if isExtentionValid {
			parts := strings.Split(extention, "_")
			extention = ""
			for _, part := range parts {
				part = strings.ToUpper(string(part[0])) + part[1:]
				extention += part
			}
			validExtentions = append(validExtentions, extention)
		}
	}
	return validExtentions
}

func GetValidFilters(initialValues string, validValues ...string) []string {
	filter := strings.Split(initialValues, ",")
	validFilters := make([]string, 0)
	for _, filter := range filter {
		filter = strings.ToLower(filter)
		isFilterValid := Contains(validValues, filter)
		if isFilterValid {
			if filter == "creation_time" {
				filter = "created_at"
			}
			validFilters = append(validFilters, filter)
		}
	}
	return validFilters
}

func Contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}