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
	database *gorm.DB
	providers oauthproviders.Providers
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		database: database.Instance,
		providers: oauthproviders.Instance,
	}
}

func (authRepo *AuthRepository) RegisterWithEmailAndPassword(fullName, email, password string) (apiError *sharedUtils.APIError) {
	database := authRepo.database

	var exist bool
	err := database.Model(&models.User{}).Select("count(*) > 0").Where("email = ?", email).Find(&exist).Error
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	if exist {
		return &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message: "email already in use",
		}
	}

	hashedPassword, err := authUtils.HashPassword(password)
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	err = database.Create(&models.User{
		FullName: fullName,
		Password: hashedPassword,
		Email: email,
	}).Error
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (authRepo *AuthRepository) LoginWithEmailAndPassword(email, password string) (result gin.H, apiError *sharedUtils.APIError) {
	database := authRepo.database

	var storedUser models.User
	err := database.Where("email = ?", email).First(&storedUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &sharedUtils.APIError{
				StatusCode: http.StatusBadGateway,
				Message: "email not found",
			}
		} else {
			return nil, &sharedUtils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	passwordMatches := authUtils.VerifyPasswordHash(password, storedUser.Password)
	if !passwordMatches {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: "wrong password",
		}
	}

	if emailVerified := storedUser.EmailVerified; !emailVerified {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: "email not verified",
		}
	}

	idToken, err := authUtils.CreateIdToken(
		storedUser.ID,
		storedUser.Email,
		storedUser.EmailVerified,
	)
	if err != nil {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	refreshToken, err := authUtils.CreateRefreshToken(storedUser.ID)
	if err != nil {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return gin.H{
		"idToken": idToken,
		"refreshToken": refreshToken,
	}, nil
}

func (authRepo *AuthRepository) Authorization(authorization string) (result gin.H,apiError *sharedUtils.APIError) {
	if authorization == "" {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message: "undefined authorization",
		}
	}
	idToken := authorization[len("Bearer "):]
	claims, err := authUtils.VerifyToken(idToken)
	if err != nil {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusUnauthorized,
			Message: err.Error(),
		}
	}

	return gin.H{
		"email":         claims["email"],
		"id":            claims["id"],
		"emailVerified": claims["emailVerified"],
		"isAdmin":       claims["isAdmin"],
		"disabled":      claims["disabled"],
		"idToken":       idToken,
	}, nil
}

func (authRepo *AuthRepository) AuthorizationWithEmailVerification(emailVerified bool) (apiError *sharedUtils.APIError) {
	if !emailVerified {
		return &sharedUtils.APIError{
			StatusCode: http.StatusUnauthorized,
		}
	}

	return nil
}

func (authRepo *AuthRepository) RefreshIdToken(authorization string) (result gin.H, apiError *sharedUtils.APIError) {
	if authorization == "" {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message: "undefined authorization",
		}
	}

	refreshToken := authorization[len("Bearer "):]
	if refreshToken == "" {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message: "undefined refresh token",
		}
	}

	claims, err := authUtils.VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusUnauthorized,
		}
	}

	database := authRepo.database

	id := uint(claims["id"].(float64))
	var idToken string
	var user models.User
	err = database.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &sharedUtils.APIError{
				StatusCode: http.StatusNotFound,
				Message:  "user not found",
			}
		}
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	idToken, err = authUtils.CreateIdToken(
		user.ID,
		user.Email,
		user.EmailVerified,
	)
	if err != nil {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return gin.H{
		"idToken": idToken,
	}, nil
}

func (authRepo *AuthRepository) SendEmailVerificationLink(toEmail string, url string) (apiError *sharedUtils.APIError) {
	idToken, err := authUtils.CreateIdTokenFromEmail(toEmail)
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	verificationLink := url + "/" + idToken
	message := fmt.Sprintf("Subject: Email verification link!\nThis is email verification link from learn_oo\n%v\nif you do not have to do with it dont browse it!", verificationLink)

	err = email.SendEmailMessage(toEmail, message)

	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (authRepo *AuthRepository) VerifyEmail(email string) (apiError *sharedUtils.APIError) {
	database := authRepo.database

	var user models.User
	err := database.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &sharedUtils.APIError{
				StatusCode: http.StatusNotFound,
				Message: "email not found",
			}
		}
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if user.EmailVerified {
		return &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message: "user already verified",
		}
	}

	err = database.Model(&models.User{}).Where("email = ?", email).Update("email_verified", true).Error
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (authRepo *AuthRepository) SendPasswordResetLink(toEmail string, url string) (apiError *sharedUtils.APIError) {
	idToken, err := authUtils.CreateIdTokenFromEmail(toEmail)
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	resetLink := url + "/" + idToken
	message := fmt.Sprintf("Subject: Password reset link!\nThis is password reset link from kinema\n%v\nif you do not have to do with it dont browse it!", resetLink)
	err = email.SendEmailMessage(toEmail, message)

	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (authRepo *AuthRepository) ResetPassword(email string, newPassword string) (error *sharedUtils.APIError) {
	database := authRepo.database

	newPasswordHash, err := authUtils.HashPassword(newPassword)
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	err = database.Model(&models.User{}).Where("email = ?", email).Update("password", newPasswordHash).Error
	if err != nil {
		return &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (authRepo *AuthRepository) OAuth(provider string, successURL string, failureURL string) (result gin.H, apiError *sharedUtils.APIError) {
	ok := oauthproviders.IsSupportedProvider(provider)
	if !ok {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message: "provider not supported",
		}
	}
	providers := authRepo.providers
	oauthConfig := providers[provider].Config

	return gin.H{
		"oauthConfig": oauthConfig,
	},nil
}

func (authRepo *AuthRepository) OAuthCallback(provider string, code string, context context.Context) (result gin.H, apiError *sharedUtils.APIError) {
	ok := oauthproviders.IsSupportedProvider(provider)
	if !ok {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusBadRequest,
			Message: "provider not supported",
		}
	}

	authProvider := authRepo.providers[provider]

	oauthConfig := authProvider.Config
	token, err := oauthConfig.Exchange(context, code)
	if err != nil {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	client := oauthConfig.Client(context, token)
	response, err := client.Get(authProvider.UserInfoURL)
	if err != nil {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer response.Body.Close()

	userData := gin.H{}
	if err := json.NewDecoder(response.Body).Decode(&userData); err != nil {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	emailVerified := true

	user := models.User{
		EmailVerified: emailVerified,
	}

	switch provider {
	case "google":
		user.FullName = userData["name"].(string)
		user.Email = userData["email"].(string)
	case "facebook":
		user.FullName = userData["name"].(string)
		user.Email = userData["email"].(string)
	}

	var database = authRepo.database

	var existingUser models.User
	err = database.Where("email = ?", user.Email).First(&existingUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = database.Create(&user).Error
			if err != nil {
				return nil, &sharedUtils.APIError{
					StatusCode: http.StatusInternalServerError,
					Message: err.Error(),
				}
			}
			existingUser = user
		} else {
			return nil, &sharedUtils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	if existingUser.Email != user.Email || existingUser.FullName != user.FullName {
		existingUser.Email = user.Email
		existingUser.FullName = user.FullName
		err = database.Save(&existingUser).Error
		if err != nil {
			return nil, &sharedUtils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	refreshToken, err := authUtils.CreateRefreshToken(existingUser.ID)
	if err != nil {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	idToken, err := authUtils.CreateIdToken(
		existingUser.ID,
		existingUser.Email,
		existingUser.EmailVerified,
	)
	if err != nil {
		return nil, &sharedUtils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return gin.H{
		"idToken": idToken,
		"refreshToken": refreshToken,
	}, nil
}