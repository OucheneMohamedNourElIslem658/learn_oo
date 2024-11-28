package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/courses/controllers"
	middlewares "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/middlewares"
)

type CoursesRouter struct {
	Router            *gin.RouterGroup
	coursesController *controllers.CoursesController
	authMiddlewares   *middlewares.AuthorizationMiddlewares
}

func NewCoursesRouter(router *gin.RouterGroup) *CoursesRouter {
	return &CoursesRouter{
		Router:          router,
		coursesController:  controllers.NewCoursesController(),
		authMiddlewares: middlewares.NewAuthorizationMiddlewares(),
	}
}

func (cr *CoursesRouter) RegisterRoutes() {
	router := cr.Router
	coursesController := cr.coursesController

	authMiddlewares := cr.authMiddlewares

	authorization := authMiddlewares.Authorization()
	authorizationWithEmailVerification := authMiddlewares.AuthorizationWithEmailVerification()
	AuthorizationWithAuthorCheck := authMiddlewares.AuthorizationWithAuthorCheck()

	router.POST("/", authorization, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, coursesController.CreateCourse)
}

