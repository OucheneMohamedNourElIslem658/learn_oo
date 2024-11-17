package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

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
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": utils.ValidationErrorResponse(err),
		})
		fmt.Println(err)
		return
	}

	authRepository := authcontroller.authRepository
	if err := authRepository.RegisterWithEmailAndPassword(body.FullName, body.Email, body.Password); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (authcontroller *AuthController) LoginWithEmailAndPassword(ctx *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,password"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": utils.ValidationErrorResponse(err),
		})
		return
	}

	authRepository := authcontroller.authRepository
	if result, err := authRepository.LoginWithEmailAndPassword(body.Email, body.Password); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
	} else {
		ctx.SetCookie("id_token", result["idToken"].(string), 3600, "/", "localhost", false, true)
		ctx.SetCookie("refresh_token", result["refreshToken"].(string), 3600, "/", "localhost", false, true)
	}
}

func (authcontroller *AuthController) SendEmailVerificationLink(ctx *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": utils.ValidationErrorResponse(err),
		})
	}

	authRepository := authcontroller.authRepository

	hostURL := "http://" + ctx.Request.Host + "/api/v1/users/auth/serve-email-verification-template"
	if err := authRepository.SendEmailVerificationLink(body.Email, hostURL); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
	} else {
		ctx.Status(http.StatusOK)
	}
}

func (authcontroller *AuthController) VerifyEmail(ctx *gin.Context) {
	authRepository := authcontroller.authRepository

	idToken := ctx.Param("id-token")
	authorization := fmt.Sprintf("Bearer %v", idToken)
	result, err := authRepository.Authorization(authorization)

	if err == nil {
		email := result["email"].(string)
		if err := authRepository.VerifyEmail(email); err != nil {
			ctx.JSON(err.StatusCode, gin.H{
				"message": err.Message,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Your email has been verified!",
		})
		return
	}

	ctx.JSON(err.StatusCode, gin.H{
		"message": err.Message,
	})
}

func (authcontroller *AuthController) ServeEmailVerificationTemplate(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "email_verification.html", nil)
}

func (authcontroller *AuthController) RefreshIdToken(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")

	repository := authcontroller.authRepository

	if result, err := repository.RefreshIdToken(authorization); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
	} else {
		ctx.Status(http.StatusOK)
		ctx.SetCookie("id_token", result["idToken"].(string), 3600, "/", "localhost", false, true)
	}
}

func (authcontroller *AuthController) SendPasswordResetLink(ctx *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": utils.ValidationErrorResponse(err),
		})
	}

	authRepository := authcontroller.authRepository

	hostURL := "http://" + ctx.Request.Host + "/api/v1/users/auth/serve-reset-password-form"
	if err := authRepository.SendPasswordResetLink(body.Email, hostURL); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Message,
		})
	} else {
		ctx.Status(http.StatusOK)
	}
}

func (authcontroller *AuthController) ResetPassword(ctx *gin.Context) {
	var body struct {
		Password string `json:"password" binding:"required,password"`
	}

	fmt.Println(body.Password)
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": utils.ValidationErrorResponse(err),
		})
		return
	}

	authRepository := authcontroller.authRepository

	idToken := ctx.Param("id-token")
	authorization := fmt.Sprintf("Bearer %v", idToken)
	if result, err := authRepository.Authorization(authorization); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Message,
		})
	} else {
		email := result["email"].(string)
		newPassword := body.Password
		if err = authRepository.ResetPassword(email, newPassword); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Message,
			})
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (authcontroller *AuthController) ServeResetPasswordForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "reset_password_form.html", nil)
}

// func (authcontroller *AuthController) OAuth(ctx *gin.Context) {
// 	provider := r.PathValue("provider")

// 	var body struct {
// 		IsAdmin bool `json:"is_admin"`
// 		SuccessURL string `json:"success_url"`
// 		FailureURL string `json:"failure_url"`
// 	}

// 	query := r.URL.Query()
// 	isAdminString := query.Get("is_admin")
// 	isAdmin, _ := strconv.ParseBool(isAdminString)
// 	body.IsAdmin = isAdmin

// 	successURL := query.Get("success_url")
// 	failureURL := query.Get("failure_url")
// 	body.SuccessURL = successURL
// 	body.FailureURL = failureURL

// 	bodyBytes, _ := json.Marshal(&body)

// 	authRepository := authcontroller.authRepository
// 	status, result := authRepository.OAuth(provider, successURL, failureURL)
// 	if status != http.StatusOK {
// 		w.WriteHeader(status)
// 		reponse, _ := json.Marshal(result)
// 		w.Write(reponse)
// 		return
// 	}

// 	oauthConfig := result["oauthConfig"].(*oauth2.Config)
// 	url := oauthConfig.AuthCodeURL(string(bodyBytes), oauth2.AccessTypeOffline)
// 	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
// }

// func (authcontroller *AuthController) OAuthCallback(ctx *gin.Context) {
// 	provider := r.PathValue("provider")

// 	query := r.URL.Query()
// 	code := query.Get("code")

// 	var metadata struct {
// 		IsAdmin bool `json:"is_admin"`
// 		SuccessURL string `json:"success_url"`
// 		FailureURL string `json:"failure_url"`
// 	}
// 	state := query.Get("state")
// 	json.Unmarshal([]byte(state), &metadata)

// 	authRepository := authcontroller.authRepository
// 	status, result := authRepository.OAuthCallback(provider, code, metadata.IsAdmin, r.Context())
// 	if status == http.StatusOK {
// 		if idToken, ok := result["id_token"].(string); ok {
// 			http.Redirect(w, r, fmt.Sprintf("%v?id_token=%v", metadata.SuccessURL, idToken), http.StatusFound)
// 		} else {
// 			err := errors.New("INTERNAL_SERVER_ERROR")
// 			http.Redirect(w, r, fmt.Sprintf("%v?error=%v", metadata.FailureURL, err.Error()), http.StatusFound)
// 		}
// 		return
// 	}

// 	http.Redirect(w, r, fmt.Sprintf("%v?error=%v", metadata.FailureURL, result["error"]), http.StatusFound)
// }
