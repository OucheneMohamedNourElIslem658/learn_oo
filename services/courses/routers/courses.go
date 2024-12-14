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
	testsController                     *controllers.TestsController
	objectivesAndRequirementsController *controllers.ObjectivesAndRequirementsController
	authMiddlewares                     *authMiddlewares.AuthorizationMiddlewares
	coursesMiddlewares                  *middlewares.CoursesMiddlewares
}

func NewCoursesRouter(router *gin.RouterGroup) *CoursesRouter {
	return &CoursesRouter{
		Router:                              router,
		coursesController:                   controllers.NewCoursesController(),
		chaptersController:                  controllers.NewChaptersController(),
		lessonsController:                   controllers.NewLessonsController(),
		testsController:                     controllers.NewTestsController(),
		objectivesAndRequirementsController: controllers.NewObjectivesAndRequirementsController(),
		authMiddlewares:                     authMiddlewares.NewAuthorizationMiddlewares(),
		coursesMiddlewares:                  middlewares.NewCoursesMiddlewares(),
	}
}

func (cr *CoursesRouter) RegisterRoutes() {
	router := cr.Router

	coursesController := cr.coursesController
	chaptersController := cr.chaptersController
	lessonsController := cr.lessonsController
	testsController := cr.testsController
	objectivesAndRequirementsController := cr.objectivesAndRequirementsController

	authMiddlewares := cr.authMiddlewares

	authorization := authMiddlewares.Authorization()
	authorizationWithIDCheck := authMiddlewares.AuthorizationWithIDCheck()
	AuthorizationWithAuthorCheck := authMiddlewares.AuthorizationWithAuthorCheck()

	coursesMiddlewares := cr.coursesMiddlewares

	checkCourseExistance := coursesMiddlewares.CheckCourseExistance()
	checkChapterExistance := coursesMiddlewares.CheckChapterExistance()
	checkTestExistance := coursesMiddlewares.CheckTestExistance()

	router.POST("/", authorization, AuthorizationWithAuthorCheck, coursesController.CreateCourse)
	router.PUT("/:course_id", authorization, AuthorizationWithAuthorCheck, coursesController.UpdateCourse)
	router.PUT("/:course_id/image", authorization, AuthorizationWithAuthorCheck, checkCourseExistance, coursesController.UpdateCourseImage)
	router.PUT("/:course_id/video", authorization, AuthorizationWithAuthorCheck, checkCourseExistance, coursesController.UpdateCourseVideo)
	router.DELETE("/:course_id", authorization, AuthorizationWithAuthorCheck, coursesController.DeleteCourse)
	router.GET("/:course_id", authorization, coursesController.GetCourse)
	router.GET("/", authorization, coursesController.GetCourses)

	objectivesRouter := router.Group("/:course_id/objectives")
	objectivesRouter.POST("/", authorization, AuthorizationWithAuthorCheck, checkCourseExistance, objectivesAndRequirementsController.CreateObjective)
	objectivesRouter.DELETE("/:objective_id", authorization, AuthorizationWithAuthorCheck, checkCourseExistance, objectivesAndRequirementsController.DeleteObjective)

	requirementsRouter := router.Group("/:course_id/requirements")
	requirementsRouter.POST("/", authorization, AuthorizationWithAuthorCheck, checkCourseExistance, objectivesAndRequirementsController.CreateRequirement)
	requirementsRouter.DELETE("/:requirement_id", authorization, AuthorizationWithAuthorCheck, checkCourseExistance, objectivesAndRequirementsController.DeleteRequirement)

	categoriesRouter := router.Group("/categories")
	categoriesRouter.POST("/", coursesController.CreateCategory)
	categoriesRouter.DELETE("/:category_id", coursesController.DeleteCategory)
	categoriesRouter.GET("/", coursesController.GetCategories)

	chaptersRouter := router.Group("/:course_id/chapters")
	chaptersRouter.POST("/", authorization, AuthorizationWithAuthorCheck, checkCourseExistance, chaptersController.CreateChapter)
	chaptersRouter.PUT("/:chapter_id", authorization, AuthorizationWithAuthorCheck, checkCourseExistance, chaptersController.UpdateChapter)
	chaptersRouter.DELETE("/:chapter_id", authorization, AuthorizationWithAuthorCheck, checkCourseExistance, chaptersController.DeleteChapter)
	chaptersRouter.GET("/:chapter_id", chaptersController.GetChapter)

	lessonsRouter := chaptersRouter.Group("/:chapter_id/lessons")
	lessonsRouter.POST("/create-with-content", authorization, AuthorizationWithAuthorCheck, checkChapterExistance, lessonsController.CreateLessonWithContent)
	lessonsRouter.POST("/create-with-video", authorization, AuthorizationWithAuthorCheck, checkChapterExistance, lessonsController.CreateLessonWithVideo)
	lessonsRouter.PUT("/:lesson_id", authorization, AuthorizationWithAuthorCheck, checkChapterExistance, lessonsController.UpdateLesson)
	lessonsRouter.PUT("/:lesson_id/video", authorization, AuthorizationWithAuthorCheck, checkChapterExistance, lessonsController.UpdateLessonVideo)
	lessonsRouter.DELETE("/:lesson_id", authorization, AuthorizationWithAuthorCheck, checkChapterExistance, lessonsController.DeleteLesson)
	lessonsRouter.GET("/:lesson_id", authorization, authorizationWithIDCheck, lessonsController.GetLesson)

	testsRouter := chaptersRouter.Group("/:chapter_id/tests")
	testsRouter.POST("/", authorization, AuthorizationWithAuthorCheck, checkChapterExistance, testsController.CreateTest)
	testsRouter.PUT("/:test_id", authorization, AuthorizationWithAuthorCheck, checkChapterExistance, testsController.UpdateTest)
	testsRouter.DELETE("/:test_id", authorization, AuthorizationWithAuthorCheck, checkChapterExistance, testsController.DeleteTest)
	testsRouter.GET("/:test_id", authorization, authorizationWithIDCheck, testsController.GetTest)

	questionsRouter := testsRouter.Group("/:test_id/questions")
	questionsRouter.POST("/", authorization, AuthorizationWithAuthorCheck, checkTestExistance, testsController.CreateQuestion)
	questionsRouter.PUT("/:question_id", authorization, AuthorizationWithAuthorCheck, checkTestExistance, testsController.UpdateQuestion)
	questionsRouter.DELETE("/:question_id", authorization, AuthorizationWithAuthorCheck, checkTestExistance, testsController.DeleteQuestion)
	questionsRouter.GET("/:question_id", authorization, authorizationWithIDCheck, testsController.GetQuestion)
}
