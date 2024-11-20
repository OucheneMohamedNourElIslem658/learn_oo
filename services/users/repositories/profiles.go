package repositories

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	gorm "gorm.io/gorm"

	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	filestorage "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/file_storage"
	models "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	utils "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
	gin "github.com/gin-gonic/gin"
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
				Message:    "user not found",
			}
		}
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
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
				Message:    "user not found",
			}
		}

		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if fullName != "" {
		existingUser.FullName = fullName
	}

	err = database.Save(&existingUser).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return &utils.APIError{
		StatusCode: http.StatusOK,
	}
}

func (pr *ProfilesRepository) UpdateUserImage(image multipart.File, userID string) (apiError *utils.APIError) {
	database := pr.database
	filestorage := pr.fileStorage

	var existingImage models.File
	err := database.Where("user_id = ?", userID).First(&existingImage).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if err == nil {
		if err := database.Where("id = ?", existingImage.ID).Unscoped().Delete(&existingImage).Error; err != nil {
			return &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
		if existingImage.ImageKitID != nil {
			if err := filestorage.DeleteFile(*existingImage.ImageKitID); err != nil {
				return &utils.APIError{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
				}
			}
		}
	}

	path := fmt.Sprintf("learn_oo/images/users/%v", userID)
	uploadData, err := filestorage.UploadFile(image, path)
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
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
			Message:    err.Error(),
		}
	}

	return &utils.APIError{
		StatusCode: http.StatusOK,
	}
}

func (pr *ProfilesRepository) UpgradeToAuthor(id uint) (apiError *utils.APIError) {
	database := pr.database
	var user models.User
	err := database.Where("id = ?", user).Preload("AuthorProfile").First(&user).Error
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
	if user.AuthorProfile == nil {
		author := models.Author{
			UserID: user.ID,
		}
		err := database.Create(&author).Error
		if err != nil {
			return &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message: err.Error(),
			}
		}

		return nil
	}

	return &utils.APIError{
		StatusCode: http.StatusBadRequest,
		Message: "user is already an author",
	}
}

func (pr *ProfilesRepository) DowngradeFromAuthor(id, userID string) (apiError *utils.APIError) {
	database := pr.database
	var author models.Author
	deleteResult := database.Where("id = ? and user_id = ?", id, userID).Unscoped().Delete(&author)
	err := deleteResult.Error
	affectedRows := deleteResult.RowsAffected
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	if affectedRows == 0 {
		return &utils.APIError{
			StatusCode: http.StatusNotFound,
			Message: "author not found",
		}
	}
	
	return nil
}

func (pr *ProfilesRepository) UpdateAuthor(id, userID string, bio gin.H) (apiError *utils.APIError) {
	database := pr.database

	var existingAuthor models.Author
	err := database.Where("id = ?", id).First(&existingAuthor).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message: "author not found",
			}
		}

		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if bio != nil {
		existingAuthor.Bio = bio
	}

	err = database.Save(&existingAuthor).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (pr *ProfilesRepository) AddAuthorAccomplishments(id string, files []multipart.File) (apiError *utils.APIError) {
	filestorage := pr.fileStorage
	uploadData, errs := filestorage.UploadFiles(files, fmt.Sprintf("/files/authors/%v", id))
	if errs != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: strings.Join(errs, " "),
		}
	}

	database := pr.database
	var accomplishments []models.File

	for _, fileUploadData := range uploadData {
		if fileUploadData != nil {
			file := models.File{
				URL:          fileUploadData.Url,
				ThumbnailURL: &fileUploadData.ThumbnailUrl,
				ImageKitID:   &fileUploadData.FileId,
				AuthorID:     &id,
			}
			accomplishments = append(accomplishments, file)
		}
	}

	err := database.Create(&accomplishments).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (pr *ProfilesRepository) DeleteAuthorAccomplishment(id uint, authorID string) (apiError *utils.APIError) {
	database := pr.database
	filestorage := pr.fileStorage

	var file models.File
	deleteResult := database.Where("id = ? and author_id = ?", id, authorID).Unscoped().Delete(&file)
	err := deleteResult.Error
	affectedRows := deleteResult.RowsAffected
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	if affectedRows == 0 {
		return &utils.APIError{
			StatusCode: http.StatusNotFound,
			Message: "file not found",
		}
	}

	if file.ImageKitID != nil {
		if err := filestorage.DeleteFile(*file.ImageKitID); err != nil {
			return &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
	}

	return nil
}

func (pr *ProfilesRepository) GetAuthor(id uint, appendWith string) (author *models.Author, apiError *utils.APIError) {
	database := pr.database
	query := database.Model(&models.Author{})

	validExtentions := utils.GetValidExtentions(
		appendWith,
		"accomplishments",
		"user",
	)
	for _, extention := range validExtentions {
		query.Preload(extention)
	}

	var existingAuthor models.Author
	err := query.Where("id = ?", id).First(&existingAuthor).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message: "author not found",
			}
		}
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &existingAuthor, nil
}
