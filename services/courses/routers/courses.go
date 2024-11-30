package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/courses/controllers"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/courses/middlewares"
	authMiddlewares "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/middlewares"
)

type CoursesRouter struct {
	Router              *gin.RouterGroup
	coursesController   *controllers.CoursesController
	chaptersControllers *controllers.ChaptersController
	authMiddlewares     *authMiddlewares.AuthorizationMiddlewares
	ChaptersMiddlewares *middlewares.ChaptersMiddlewares
}

func NewCoursesRouter(router *gin.RouterGroup) *CoursesRouter {
	return &CoursesRouter{
		Router:              router,
		coursesController:   controllers.NewCoursesController(),
		authMiddlewares:     authMiddlewares.NewAuthorizationMiddlewares(),
		ChaptersMiddlewares: middlewares.NewChaptersMiddlewares(),
	}
}

func (cr *CoursesRouter) RegisterRoutes() {
	router := cr.Router
	coursesController := cr.coursesController
	chaptersController := cr.chaptersControllers

	authMiddlewares := cr.authMiddlewares

	authorization := authMiddlewares.Authorization()
	authorizationWithEmailVerification := authMiddlewares.AuthorizationWithEmailVerification()
	AuthorizationWithAuthorCheck := authMiddlewares.AuthorizationWithAuthorCheck()

	chaptersMiddlewares := cr.ChaptersMiddlewares

	CheckCourseExistance := chaptersMiddlewares.CheckCourseExistance()

	router.POST("/", authorization, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, coursesController.CreateCourse)
	router.PUT("/:id", authorization, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, coursesController.UpdateCourse)
	router.DELETE("/:id", authorization, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, coursesController.DeleteCourse)
	router.GET("/:id", coursesController.GetCourse)
	router.GET("/", coursesController.GetCourses)

	categoriesRouter := router.Group("/categories")
	categoriesRouter.POST("/", coursesController.CreateCategory)
	categoriesRouter.DELETE("/:id", coursesController.DeleteCategory)
	categoriesRouter.GET("/", coursesController.GetCategories)

	chaptersRouter := router.Group("/:course_id/chapters")
	chaptersRouter.POST("/", authorization, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, CheckCourseExistance, chaptersController.CreateChapter)
	chaptersRouter.PUT("/:id", authorization, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, CheckCourseExistance, chaptersController.UpdateChapter)
	chaptersRouter.DELETE("/:id", authorization, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, CheckCourseExistance, chaptersController.DeleteChapter)
	chaptersRouter.GET("/:id", chaptersController.GetChapter)
}
