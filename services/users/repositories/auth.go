package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	gin "github.com/gin-gonic/gin"
	gorm "gorm.io/gorm"

	authUtils "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/utils"
	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	email "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/email"
	models "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	oauthproviders "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/oauth_providers"
	sharedUtils "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
)

type AuthRepository struct {
	database  *gorm.DB
	providers oauthproviders.Providers
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		database:  database.Instance,
		providers: oauthproviders.Instance,
	}
}

func (ar *AuthRepository) RegisterWithEmailAndPassword(fullName, email, password string) (apiError *sharedUtils.APIError) {
	database := ar.database

	var exist bool
	err := database.Model(&models.User{}).Select("count(*) > 0").Where("email = ?", email).Find(&exist).Error
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	if exist {
		return &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    "email already in use",
		}
	}

	hashedPassword, err := authUtils.HashPassword(password)
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	err = database.Create(&models.User{
		FullName: fullName,
		Password: hashedPassword,
		Email:    email,
	}).Error
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	return nil
}

func (ar *AuthRepository) LoginWithEmailAndPassword(email, password string) (idToken, refreshToken *string, apiError *sharedUtils.APIError) {
	database := ar.database

	var storedUser models.User
	err := database.Where("email = ?", email).Preload("AuthorProfile").First(&storedUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, &sharedUtils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "email not found",
			}
		} else {
			return nil, nil, &sharedUtils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
	}

	passwordMatches := authUtils.VerifyPasswordHash(password, storedUser.Password)
	if !passwordMatches {
		return nil, nil, &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    "wrong password",
		}
	}

	if emailVerified := storedUser.EmailVerified; !emailVerified {
		return nil, nil, &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    "email not verified",
		}
	}

	author := storedUser.AuthorProfile
    authorID := func() *string {
		if author == nil {
			return nil
		}
		return &author.ID
	}()

	createdIdToken, err := authUtils.CreateIdToken(
		storedUser.ID,
		authorID,
		storedUser.EmailVerified,
	)
	if err != nil {
		return nil, nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	createdRefreshToken, err := authUtils.CreateRefreshToken(storedUser.ID)
	if err != nil {
		return nil, nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return &createdIdToken, &createdRefreshToken, nil
}

func (ar *AuthRepository) AuthorizationWithEmailVerification(emailVerified bool) (apiError *sharedUtils.APIError) {
	if !emailVerified {
		return &sharedUtils.APIError{
			StatusCode: http.StatusUnauthorized,
		}
	}

	return nil
}

func (ar *AuthRepository) RefreshIdToken(authorization string) (idToken *string, apiError *sharedUtils.APIError) {
	if authorization == "" {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    "undefined authorization",
		}
	}

	refreshToken := authorization[len("Bearer "):]
	if refreshToken == "" {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    "undefined refresh token",
		}
	}

	claims, err := authUtils.VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusUnauthorized,
			Message: "refresh token expired",
		}
	}

	database := ar.database

	id, ok := claims["id"].(string)
	if !ok {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: "casting id failed",
		}
	}

	var user models.User
	err = database.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &sharedUtils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "user not found",
			}
		}
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}

	author := user.AuthorProfile
    authorID := func() *string {
		if author == nil {
			return nil
		}
		return &author.ID
	}()

	createdIDToken, err := authUtils.CreateIdToken(
		user.ID,
		authorID,
		user.EmailVerified,
	)
	if err != nil {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}

	return &createdIDToken, nil
}

func (ar *AuthRepository) SendEmailVerificationLink(toEmail string, url string) (apiError *sharedUtils.APIError) {
	idToken, err := authUtils.CreateIdTokenFromEmail(toEmail)
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	verificationLink := url + "/" + idToken
	message := fmt.Sprintf("Subject: Email verification link!\nThis is email verification link from learn_oo\n%v\nif you do not have to do with it dont browse it!", verificationLink)

	err = email.SendEmailMessage(toEmail, message)

	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (ar *AuthRepository) VerifyEmail(email string) (apiError *sharedUtils.APIError) {
	database := ar.database

	var user models.User
	err := database.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &sharedUtils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "email not found",
			}
		}
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if user.EmailVerified {
		return &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    "user already verified",
		}
	}

	err = database.Model(&models.User{}).Where("email = ?", email).Update("email_verified", true).Error
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (ar *AuthRepository) SendPasswordResetLink(toEmail string, url string) (apiError *sharedUtils.APIError) {
	idToken, err := authUtils.CreateIdTokenFromEmail(toEmail)
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	resetLink := url + "/" + idToken
	message := fmt.Sprintf("Subject: Password reset link!\nThis is password reset link from kinema\n%v\nif you do not have to do with it dont browse it!", resetLink)
	err = email.SendEmailMessage(toEmail, message)

	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (ar *AuthRepository) ResetPassword(email string, newPassword string) (error *sharedUtils.APIError) {
	database := ar.database

	newPasswordHash, err := authUtils.HashPassword(newPassword)
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	err = database.Model(&models.User{}).Where("email = ?", email).Update("password", newPasswordHash).Error
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (ar *AuthRepository) OAuth(provider string, successURL string, failureURL string) (result gin.H, apiError *sharedUtils.APIError) {
	ok := oauthproviders.IsSupportedProvider(provider)
	if !ok {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    "provider not supported",
		}
	}
	providers := ar.providers
	oauthConfig := providers[provider].Config

	return gin.H{
		"oauthConfig": oauthConfig,
	}, nil
}

func (ar *AuthRepository) OAuthCallback(provider string, code string, context context.Context) (idToken, refreshToken *string, apiError *sharedUtils.APIError) {
	ok := oauthproviders.IsSupportedProvider(provider)
	if !ok {
		return nil, nil, &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    "provider not supported",
		}
	}

	authProvider := ar.providers[provider]

	oauthConfig := authProvider.Config
	token, err := oauthConfig.Exchange(context, code)
	if err != nil {
		return  nil, nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	client := oauthConfig.Client(context, token)
	response, err := client.Get(authProvider.UserInfoURL)
	if err != nil {
		return  nil, nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	defer response.Body.Close()

	userData := gin.H{}
	if err := json.NewDecoder(response.Body).Decode(&userData); err != nil {
		return  nil, nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	emailVerified := true

	user := models.User{
		EmailVerified: emailVerified,
	}

	// Handle the profile pic url:

	switch provider {
	case "google":
		user.FullName = userData["name"].(string)
		user.Email = userData["email"].(string)
		user.Image = &models.File{
			URL: userData["picture"].(string),
		}
	case "facebook":
		user.FullName = userData["name"].(string)
		user.Email = userData["email"].(string)
		user.Image = &models.File{
			URL: userData["picture"].(map[string]interface{})["data"].(map[string]interface{})["url"].(string),
		}
	}

	var database = ar.database

	var existingUser models.User
	err = database.Where("email = ?", user.Email).Preload("Image").First(&existingUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = database.Create(&user).Error
			if err != nil {
				return  nil, nil, &sharedUtils.APIError{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
				}
			}
			existingUser = user
		} else {
			return  nil, nil, &sharedUtils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
	}

	existingUser.Email = user.Email
	existingUser.FullName = user.FullName
	if user.Image == nil {
		existingUser.Image = user.Image
	}
	err = database.Session(&gorm.Session{FullSaveAssociations: true}).Save(&existingUser).Error
	if err != nil {
		return  nil, nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	createdRefreshToken, err := authUtils.CreateRefreshToken(existingUser.ID)
	if err != nil {
		return  nil, nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	author := existingUser.AuthorProfile
    authorID := func() *string {
		if author == nil {
			return nil
		}
		return &author.ID
	}()

	createdIDToken, err := authUtils.CreateIdToken(
		existingUser.ID,
		authorID,
		existingUser.EmailVerified,
	)
	if err != nil {
		return  nil, nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return &createdIDToken, &createdRefreshToken, nil
}
