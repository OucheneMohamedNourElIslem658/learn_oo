package controllers

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"

	repositories "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/repositories"
	utils "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
)

type ProfilesController struct {
	profilesRepository *repositories.ProfilesRepository
}

func NewProfilesController() *ProfilesController {
	return &ProfilesController{
		profilesRepository: repositories.NewProfilesRepository(),
	}
}

func (pc *ProfilesController) GetUser(ctx *gin.Context) {
	id := ctx.GetString("id")
	appendWith := ctx.Query("append_with")

	profilesRepository := pc.profilesRepository
	user, err := profilesRepository.GetUser(id, appendWith)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (pc *ProfilesController) UpdateUser(ctx *gin.Context) {
	id := ctx.GetString("id")
	var body struct {
		FullName string `json:"full_name"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, message)
		fmt.Println(err.Error())
		return
	}

	profilesRepository := pc.profilesRepository
	err := profilesRepository.UpdateUser(id, body.FullName)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (pc *ProfilesController) UpdateUserImage(ctx *gin.Context) {
	id := ctx.GetString("id")
	image, imageHeader, err := ctx.Request.FormFile("image")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if imageHeader == nil || image == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "image not provided",
		})
		return
	}

	if !utils.IsImage(*imageHeader) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "the file is not an image",
		})
		return
	}

	profilesRepository := pc.profilesRepository

	apiError := profilesRepository.UpdateUserImage(id, image)
	if apiError != nil {
		ctx.JSON(apiError.StatusCode, gin.H{
			"error": apiError.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (pc *ProfilesController) UpgradeToAuthor(ctx *gin.Context) {
	id := ctx.GetString("id")

	profilesRepository := pc.profilesRepository
	err := profilesRepository.UpgradeToAuthor(id)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (pc *ProfilesController) DowngradeFromAuthor(ctx *gin.Context) {
	authorID := ctx.GetString("author_id")

	profilesRepository := pc.profilesRepository
	err := profilesRepository.DowngradeFromAuthor(authorID)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (pc *ProfilesController) UpdateAuthor(ctx *gin.Context) {
	authorID := ctx.GetString("author_id")

	var body struct {
		Bio gin.H `json:"bio"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, message)
		fmt.Println(err.Error())
		return
	}

	profilesRepository := pc.profilesRepository
	err := profilesRepository.UpdateAuthor(authorID, body.Bio)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (pc *ProfilesController) AddAuthorAccomplishments(ctx *gin.Context) {
	authorID := ctx.GetString("author_id")

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	filesHeaders := form.File["accomplishments"]

	var files []multipart.File
	for _, fileHeader := range filesHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		files = append(files, file)
	}

	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "accomplishments can not be empty",
		})
		return
	}

	profilesRepository := pc.profilesRepository
	apiError := profilesRepository.AddAuthorAccomplishments(authorID, files)
	if apiError != nil {
		ctx.JSON(apiError.StatusCode, gin.H{
			"error": apiError.Message,
		})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (pc *ProfilesController) DeleteAuthorAccomplishment(ctx *gin.Context) {
	authorID := ctx.GetString("author_id")
	fileID := ctx.Param("file_id")

	profilesRepository := pc.profilesRepository
	err := profilesRepository.DeleteAuthorAccomplishment(authorID, fileID)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (pc *ProfilesController) GetAuthor(ctx *gin.Context) {
	authorID := ctx.GetString("author_id")
	appendWith := ctx.Query("append_with")

	profilesRepository := pc.profilesRepository
	user, err := profilesRepository.GetAuthor(authorID, appendWith)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
