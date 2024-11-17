package routers

import (
	"github.com/gin-gonic/gin"

	controllers "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/controllers"
)

type AuthRouter struct {
	Router         *gin.RouterGroup
	authController *controllers.AuthController
}

func NewAuthRouter(router *gin.RouterGroup) *AuthRouter {
	return &AuthRouter{
		Router:         router,
		authController: controllers.NewAuthController(),
	}
}

func (authRouter *AuthRouter) RegisterRoutes() {
	router := authRouter.Router
	authController := authRouter.authController

	router.POST("/register-with-email-and-password", authController.RegisterWithEmailAndPassword)
	router.POST("/login-with-email-and-password", authController.LoginWithEmailAndPassword)
	router.POST("/send-email-verification-link", authController.SendEmailVerificationLink)
	router.GET("/serve-email-verification-template/:id-token", authController.ServeEmailVerificationTemplate)
	router.GET("/verify-email/:id-token", authController.VerifyEmail)
	router.GET("/refresh-id-token", authController.RefreshIdToken)
	router.POST("/send-password-reset-link", authController.SendPasswordResetLink)
	router.GET("/serve-reset-password-form/:id-token", authController.ServeResetPasswordForm)
	router.POST("/reset-password/:id-token", authController.ResetPassword)
	router.GET("/oauth/:provider/login", authController.OAuth)
	router.GET("/oauth/:provider/callback", authController.OAuthCallback)
}
