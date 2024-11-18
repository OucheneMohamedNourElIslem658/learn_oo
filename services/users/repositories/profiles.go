package repositories

import (
	"fmt"
	"mime/multipart"
	"net/http"

	gorm "gorm.io/gorm"

	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	filestorage "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/file_storage"
	models "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	utils "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
)

type ProfilesRepository struct {
	database    *gorm.DB
	fileStorage *filestorage.FileStorage
}

func NewProfilesRepository() *ProfilesRepository {
	return &ProfilesRepository{
		database:    database.Instance,
		fileStorage: filestorage.NewFileStorage(),
	}
}

func (pr *ProfilesRepository) GetUser(id, appendWith string) (user *models.User, apiError *utils.APIError) {
	database := pr.database

	query := database.Model(&models.User{})

	validExtentions := utils.GetValidExtentions(
		appendWith,
		"image",
		"author_profile",
	)
	for _, extention := range validExtentions {
		query.Preload(extention)
	}

	var existingUser models.User
	err := query.Where("id = ?", id).First(&existingUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message: "user not found",
			}
		}
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &existingUser, nil
}

func (pr *ProfilesRepository) UpdateUser(id, fullName string) (apiError *utils.APIError) {
	database := pr.database

	var existingUser models.User
	err := database.Where("id = ?", id).First(&existingUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message: "user not found",
			}
		}

		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if fullName != "" {
		existingUser.FullName = fullName
	}

	err = database.Save(&existingUser).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &utils.APIError{
		StatusCode: http.StatusOK,
	}
}

func (pr ProfilesRepository) UpdateUserImage(image multipart.File, userID string) (apiError *utils.APIError) {
	database := pr.database
	filestorage := pr.fileStorage

	var existingImage models.File
	err := database.Where("user_id = ?", userID).First(&existingImage).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if err == nil {
		if err := database.Where("id = ?", existingImage.ID).Unscoped().Delete(&existingImage).Error; err != nil {
			return &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
		if existingImage.ImageKitID != nil {
			if err := filestorage.DeleteFile(*existingImage.ImageKitID); err != nil {
				return &utils.APIError{
					StatusCode: http.StatusInternalServerError,
					Message: err.Error(),
				}
			}
		}
	}

	path := fmt.Sprintf("learn_oo/images/users/%v", userID)
	uploadData, err := filestorage.UploadFile(image, path)
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	newImage := &models.File{
		URL:          uploadData.Url,
		ImageKitID:   &uploadData.FileId,
		ThumbnailURL: &uploadData.ThumbnailUrl,
		UserID:       &userID,
	}
	if err := database.Create(&newImage).Error; err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &utils.APIError{
		StatusCode: http.StatusOK,
	}
}
