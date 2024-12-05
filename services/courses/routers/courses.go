package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/courses/controllers"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/courses/middlewares"
	authMiddlewares "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/middlewares"
)

type CoursesRouter struct {
	Router                              *gin.RouterGroup
	coursesController                   *controllers.CoursesController
	chaptersController                  *controllers.ChaptersController
	lessonsController                   *controllers.LessonsController
	objectivesAndRequirementsController *controllers.ObjectivesAndRequirementsController
	authMiddlewares                     *authMiddlewares.AuthorizationMiddlewares
	ChaptersMiddlewares                 *middlewares.ChaptersMiddlewares
}

func NewCoursesRouter(router *gin.RouterGroup) *CoursesRouter {
	return &CoursesRouter{
		Router:                              router,
		coursesController:                   controllers.NewCoursesController(),
		chaptersController:                  controllers.NewChaptersController(),
		lessonsController:                   controllers.NewLessonsController(),
		objectivesAndRequirementsController: controllers.NewObjectivesAndRequirementsController(),
		authMiddlewares:                     authMiddlewares.NewAuthorizationMiddlewares(),
		ChaptersMiddlewares:                 middlewares.NewChaptersMiddlewares(),
	}
}

func (cr *CoursesRouter) RegisterRoutes() {
	router := cr.Router

	coursesController := cr.coursesController
	chaptersController := cr.chaptersController
	lessonsController := cr.lessonsController
	objectivesAndRequirementsController := cr.objectivesAndRequirementsController

	authMiddlewares := cr.authMiddlewares

	authorization := authMiddlewares.Authorization()
	authorizationWithIDCheck := authMiddlewares.AuthorizationWithIDCheck()
	authorizationWithEmailVerification := authMiddlewares.AuthorizationWithEmailVerification()
	AuthorizationWithAuthorCheck := authMiddlewares.AuthorizationWithAuthorCheck()

	chaptersMiddlewares := cr.ChaptersMiddlewares

	CheckCourseExistance := chaptersMiddlewares.CheckCourseExistance()
	checkChapterExistance := chaptersMiddlewares.CheckChapterExistance()

	router.POST("/", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, coursesController.CreateCourse)
	router.PUT("/:course_id", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, coursesController.UpdateCourse)
	router.PUT("/:course_id/image", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, CheckCourseExistance, coursesController.UpdateCourseImage)
	router.PUT("/:course_id/video", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, CheckCourseExistance, coursesController.UpdateCourseVideo)
	router.DELETE("/:course_id", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, coursesController.DeleteCourse)
	router.GET("/:course_id", authorization, authorizationWithIDCheck, coursesController.GetCourse)
	router.GET("/", authorization, coursesController.GetCourses)

	objectivesRouter := router.Group("/:course_id/objectives")
	objectivesRouter.POST("/", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, CheckCourseExistance, objectivesAndRequirementsController.CreateObjective)
	objectivesRouter.DELETE("/:objective_id", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, CheckCourseExistance, objectivesAndRequirementsController.DeleteObjective)

	requirementsRouter := router.Group("/:course_id/requirements")
	requirementsRouter.POST("/", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, CheckCourseExistance, objectivesAndRequirementsController.CreateRequirement)
	requirementsRouter.DELETE("/:requirement_id", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, CheckCourseExistance, objectivesAndRequirementsController.DeleteRequirement)

	categoriesRouter := router.Group("/categories")
	categoriesRouter.POST("/", coursesController.CreateCategory)
	categoriesRouter.DELETE("/:category_id", coursesController.DeleteCategory)
	categoriesRouter.GET("/", coursesController.GetCategories)

	chaptersRouter := router.Group("/:course_id/chapters")
	chaptersRouter.POST("/", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, CheckCourseExistance, chaptersController.CreateChapter)
	chaptersRouter.PUT("/:chapter_id", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, CheckCourseExistance, chaptersController.UpdateChapter)
	chaptersRouter.DELETE("/:chapter_id", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, CheckCourseExistance, chaptersController.DeleteChapter)
	chaptersRouter.GET("/:chapter_id", chaptersController.GetChapter)

	lessonsRouter := chaptersRouter.Group("/:chapter_id/lessons")
	lessonsRouter.POST("/create-with-content", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, checkChapterExistance, lessonsController.CreateLessonWithContent)
	lessonsRouter.POST("/create-with-video", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, checkChapterExistance, lessonsController.CreateLessonWithVideo)
	lessonsRouter.PUT("/:lesson_id", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, checkChapterExistance, lessonsController.UpdateLesson)
	lessonsRouter.PUT("/:lesson_id/video", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, checkChapterExistance, lessonsController.UpdateLessonVideo)
	lessonsRouter.DELETE("/:lesson_id", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, AuthorizationWithAuthorCheck, checkChapterExistance, lessonsController.DeleteLesson)
	lessonsRouter.GET("/:lesson_id", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, lessonsController.GetLesson)
}
