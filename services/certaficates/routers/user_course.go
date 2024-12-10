package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/certaficates/controllers"
	authMiddlewares "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/middlewares"
)

type UserCourseRouter struct {
	Router                              *gin.RouterGroup
	userCourseController                   *controllers.UserCourseController
	authMiddlewares                     *authMiddlewares.AuthorizationMiddlewares
}

func NewUserCourseRouter(router *gin.RouterGroup) *UserCourseRouter {
	return &UserCourseRouter{
		Router:                              router,
		userCourseController:                   controllers.NewUserCourseController(),
		authMiddlewares:                     authMiddlewares.NewAuthorizationMiddlewares(),
	}
}

func (ucr *UserCourseRouter) RegisterRoutes() {
	router := ucr.Router

	userCourseController := ucr.userCourseController

	authMiddlewares := ucr.authMiddlewares

	authorization := authMiddlewares.Authorization()
	authorizationWithIDCheck := authMiddlewares.AuthorizationWithIDCheck()
	authorizationWithEmailVerification := authMiddlewares.AuthorizationWithEmailVerification()

	router.POST("/start-course/:course_id", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, userCourseController.StartCourse)
	router.POST("/pay-for-course", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, userCourseController.PayForCourse)
}