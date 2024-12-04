package routers

import (
	"github.com/gin-gonic/gin"

	controllers "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/controllers"
	middlewares "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/middlewares"
)

type AuthRouter struct {
	Router          *gin.RouterGroup
	authController  *controllers.AuthController
	authMiddlewares *middlewares.AuthorizationMiddlewares
}

func NewAuthRouter(router *gin.RouterGroup) *AuthRouter {
	return &AuthRouter{
		Router:          router,
		authController:  controllers.NewAuthController(),
		authMiddlewares: middlewares.NewAuthorizationMiddlewares(),
	}
}

func (ar *AuthRouter) RegisterRoutes() {
	router := ar.Router
	authController := ar.authController

	authMiddlewares := ar.authMiddlewares
	authorizationWithEmailCheck := authMiddlewares.AuthorizationWithEmailCheck()

	router.POST("/register-with-email-and-password", authController.RegisterWithEmailAndPassword)
	router.POST("/login-with-email-and-password", authController.LoginWithEmailAndPassword)
	router.POST("/send-email-verification-link", authController.SendEmailVerificationLink)
	router.GET("/serve-email-verification-template/:id_token", authController.ServeEmailVerificationTemplate)
	router.GET("/verify-email/:id_token", authorizationWithEmailCheck, authController.VerifyEmail)
	router.POST("/send-password-reset-link", authController.SendPasswordResetLink)
	router.GET("/serve-reset-password-form/:id_token", authController.ServeResetPasswordForm)
	router.POST("/reset-password/:id_token", authorizationWithEmailCheck, authController.ResetPassword)
	router.GET("/oauth/:provider/login", authController.OAuth)
	router.GET("/oauth/:provider/callback", authController.OAuthCallback)
}
