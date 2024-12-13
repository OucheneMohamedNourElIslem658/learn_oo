package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	repositories "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/repositories"
	utils "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
)

type AuthController struct {
	authRepository *repositories.AuthRepository
}

func NewAuthController() *AuthController {
	return &AuthController{
		authRepository: repositories.NewAuthRepository(),
	}
}

func (authcontroller *AuthController) RegisterWithEmailAndPassword(ctx *gin.Context) {
	var body struct {
		FullName string `json:"full_name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,password"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	authRepository := authcontroller.authRepository
	if err := authRepository.RegisterWithEmailAndPassword(body.FullName, body.Email, body.Password); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (authcontroller *AuthController) LoginWithEmailAndPassword(ctx *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,password"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	authRepository := authcontroller.authRepository
	if idToken, refreshToken, err := authRepository.LoginWithEmailAndPassword(body.Email, body.Password); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
	} else {
		ctx.SetCookie("id_token", *idToken, 3600, "/", "", false, true)
		ctx.SetCookie("refresh_token", *refreshToken, 3600, "/", "", false, true)
		ctx.JSON(http.StatusOK, nil)
	}
}

func (authcontroller *AuthController) SendEmailVerificationLink(ctx *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadGateway, message)
		return
	}

	authRepository := authcontroller.authRepository

	hostURL := "http://" + ctx.Request.Host + "/api/v1/users/auth/serve-email-verification-template"
	if err := authRepository.SendEmailVerificationLink(body.Email, hostURL); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, nil)
	}
}

func (authcontroller *AuthController) VerifyEmail(ctx *gin.Context) {
	authRepository := authcontroller.authRepository

	email := ctx.GetString("email")

	if err := authRepository.VerifyEmail(email); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your email has been verified!",
	})
}

func (authcontroller *AuthController) ServeEmailVerificationTemplate(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "email_verification.html", nil)
}

func (authcontroller *AuthController) RefreshIdToken(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")

	repository := authcontroller.authRepository

	if idToken, err := repository.RefreshIdToken(authorization); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, nil)
		ctx.SetCookie("id_token", *idToken, 3600, "/", "localhost", false, true)
	}
}

// func (authcontroller *AuthController) GetUserData(ctx *gin.Context) {
// 	authorization := ctx.GetHeader("Authorization")

// 	repository := authcontroller.authRepository

// 	if idToken, err := repository.RefreshIdToken(authorization); err != nil {
// 		ctx.JSON(err.StatusCode, gin.H{
// 			"error": err.Message,
// 		})
// 	} else {
// 		ctx.JSON(http.StatusOK, nil)
// 		ctx.SetCookie("id_token", *idToken, 3600, "/", "localhost", false, true)
// 	}
// }

func (authcontroller *AuthController) SendPasswordResetLink(ctx *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadGateway, message)
	}

	authRepository := authcontroller.authRepository

	hostURL := "http://" + ctx.Request.Host + "/api/v1/users/auth/serve-reset-password-form"
	if err := authRepository.SendPasswordResetLink(body.Email, hostURL); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, nil)
	}
}

func (authcontroller *AuthController) ResetPassword(ctx *gin.Context) {
	var body struct {
		Password string `json:"password" binding:"required,password"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	email := ctx.GetString("email")

	authRepository := authcontroller.authRepository

	newPassword := body.Password
	if err := authRepository.ResetPassword(email, newPassword); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (authcontroller *AuthController) ServeResetPasswordForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "reset_password_form.html", nil)
}

func (authcontroller *AuthController) OAuth(ctx *gin.Context) {
	var query struct {
		SuccessURL string `form:"success_url" json:"success_url" binding:"required"`
		FailureURL string `form:"failure_url" json:"failure_url" binding:"required"`
	}

	if err := ctx.ShouldBindQuery(&query); err != nil {
		message := utils.ValidationErrorResponse(err)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	provider := ctx.Param("provider")

	authRepository := authcontroller.authRepository
	if result, err := authRepository.OAuth(provider, query.SuccessURL, query.FailureURL); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Message,
		})
	} else {
		oauthConfig := result["oauthConfig"].(*oauth2.Config)
		queryBytes, _ := json.Marshal(&query)
		url := oauthConfig.AuthCodeURL(string(queryBytes), oauth2.AccessTypeOffline)
		ctx.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func (authcontroller *AuthController) OAuthCallback(ctx *gin.Context) {
	provider := ctx.Param("provider")

	var metadata struct {
		SuccessURL string `json:"success_url"`
		FailureURL string `json:"failure_url"`
	}
	code := ctx.Query("code")
	state := ctx.Query("state")
	json.Unmarshal([]byte(state), &metadata)

	authRepository := authcontroller.authRepository

	if idToken, refreshToken, err := authRepository.OAuthCallback(provider, code, ctx.Request.Context()); err != nil {
		failureURL := fmt.Sprintf("%v?message=%v", metadata.FailureURL, err.Message)
		ctx.Redirect(http.StatusTemporaryRedirect, failureURL)
	} else {
		ctx.SetCookie("id_token", *idToken, 3600, "/", "", false, true)
		ctx.SetCookie("refresh_token", *refreshToken, 3600, "/", "", false, true)
		ctx.Redirect(http.StatusTemporaryRedirect, metadata.SuccessURL)
	}
}
