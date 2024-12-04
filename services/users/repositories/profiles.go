package repositories

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	gorm "gorm.io/gorm"
	clause "gorm.io/gorm/clause"

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

func (UsersRouter *ProfilesRepository) GetUser(id, appendWith string) (user *models.User, apiError *utils.APIError) {
	database := UsersRouter.database

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

func (UsersRouter *ProfilesRepository) UpdateUser(id, fullName string) (apiError *utils.APIError) {
	database := UsersRouter.database

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

	return nil
}

func (UsersRouter *ProfilesRepository) UpdateUserImage(id string, image multipart.File) (apiError *utils.APIError) {
	database := UsersRouter.database
	filestorage := UsersRouter.fileStorage

	var existingImage models.File
	err := database.Where("user_id = ?", id).First(&existingImage).Error
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

	path := fmt.Sprintf("learn_oo/images/users/%v", id)
	uploadData, err := filestorage.UploadFile(image, path)
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	fmt.Println(uploadData.Url)
	newImage := models.File{
		URL:          uploadData.Url,
		ImageKitID:   &uploadData.FileId,
		ThumbnailURL: &uploadData.ThumbnailUrl,
		UserID:       &id,
	}
	if err := database.Create(&newImage).Error; err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (UsersRouter *ProfilesRepository) UpgradeToAuthor(id string) (apiError *utils.APIError) {
	database := UsersRouter.database
	var user models.User

	err := database.Where("id = ?", id).Preload("AuthorProfile").First(&user).Error
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

	if user.AuthorProfile != nil {
		return &utils.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    "user is already an author",
		}
	}

	var author models.Author
	err = database.Unscoped().Where("user_id = ?", user.ID).First(&author).Error
	if err == nil {
		err = database.Model(models.Author{}).Where("id = ?", author.ID).Unscoped().Update("deleted_at", nil).Error
		if err != nil {
			return &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}

		if err := database.Model(&models.File{}).Where("author_id = ?", author.ID).Unscoped().Update("deleted_at", nil).Error; err != nil {
			return &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}

		if err := database.Model(&models.Course{}).Where("author_id = ?", author.ID).Unscoped().Update("deleted_at", nil).Error; err != nil {
			return &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
	} else if err != gorm.ErrRecordNotFound {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	author = models.Author{UserID: user.ID}
	err = database.Create(&author).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (UsersRouter *ProfilesRepository) DowngradeFromAuthor(authorID string) (apiError *utils.APIError) {
	database := UsersRouter.database
	var author models.Author
	deleteResult := database.Where("id = ?", authorID).Select(clause.Associations).Delete(&author)
	err := deleteResult.Error
	affectedRows := deleteResult.RowsAffected
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	if affectedRows == 0 {
		return &utils.APIError{
			StatusCode: http.StatusNotFound,
			Message:    "author not found",
		}
	}

	return nil
}

func (UsersRouter *ProfilesRepository) UpdateAuthor(id string, bio gin.H) (apiError *utils.APIError) {
	database := UsersRouter.database

	var existingAuthor models.Author
	err := database.Where("id = ?", id).First(&existingAuthor).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "author not found",
			}
		}

		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if bio != nil {
		existingAuthor.Bio = bio
	}

	err = database.Save(&existingAuthor).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (UsersRouter *ProfilesRepository) AddAuthorAccomplishments(authorID string, files []multipart.File) (apiError *utils.APIError) {
	filestorage := UsersRouter.fileStorage
	uploadData, errs := filestorage.UploadFiles(files, fmt.Sprintf("/learn_oo/authors/accomplisments/%v", authorID))
	if errs != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    strings.Join(errs, " "),
		}
	}

	database := UsersRouter.database
	var accomplishments []models.File

	for _, fileUploadData := range uploadData {
		if fileUploadData != nil {
			file := models.File{
				URL:          fileUploadData.Url,
				ThumbnailURL: &fileUploadData.ThumbnailUrl,
				ImageKitID:   &fileUploadData.FileId,
				AuthorID:     &authorID,
			}
			accomplishments = append(accomplishments, file)
		}
	}

	err := database.Create(&accomplishments).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (pr *ProfilesRepository) DeleteAuthorAccomplishment(authorID, fileID string) (apiError *utils.APIError) {
	database := pr.database
	filestorage := pr.fileStorage

	var exitingFile models.File
	err := database.Where("id = ? and author_id = ?", fileID, authorID).First(&exitingFile).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "file not found",
			}
		}

		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	err = database.Delete(&exitingFile).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if exitingFile.ImageKitID != nil {
		if err := filestorage.DeleteFile(*exitingFile.ImageKitID); err != nil {
			return &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
	}

	return nil
}

func (UsersRouter *ProfilesRepository) GetAuthor(authorID string, appendWith string) (author *models.Author, apiError *utils.APIError) {
	database := UsersRouter.database
	query := database.Model(&models.Author{})

	validExtentions := utils.GetValidExtentions(
		appendWith,
		"accomplishments",
		"user",
		"courses",
	)
	for _, extention := range validExtentions {
		query.Preload(extention)
	}

	var existingAuthor models.Author
	err := query.Where("id = ?", authorID).First(&existingAuthor).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "author not found",
			}
		}
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return &existingAuthor, nil
}
