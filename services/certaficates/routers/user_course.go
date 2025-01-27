package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/certaficates/controllers"
	authMiddlewares "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/middlewares"
)

type UserCourseRouter struct {
	Router                 *gin.RouterGroup
	userCourseController   *controllers.UserCourseController
	userProgressController *controllers.UserProgressController // Add UserProgressController
	authMiddlewares        *authMiddlewares.AuthorizationMiddlewares
}

func NewUserCourseRouter(router *gin.RouterGroup) *UserCourseRouter {
	return &UserCourseRouter{
		Router:                 router,
		userCourseController:   controllers.NewUserCourseController(),
		userProgressController: controllers.NewUserProgressController(), // Initialize UserProgressController
		authMiddlewares:        authMiddlewares.NewAuthorizationMiddlewares(),
	}
}

func (ucr *UserCourseRouter) RegisterRoutes() {
	router := ucr.Router

	userCourseController := ucr.userCourseController
	userProgressController := ucr.userProgressController // Get the initialized UserProgressController

	authMiddlewares := ucr.authMiddlewares

	authorization := authMiddlewares.Authorization()
	authorizationWithIDCheck := authMiddlewares.AuthorizationWithIDCheck()

	// UserCourseController routes
	router.POST("/start-course/:course_id", authorization, authorizationWithIDCheck, userCourseController.StartCourse)
	router.POST("/pay-for-course", userCourseController.PayForCourse)

	// UserProgressController routes
	router.GET("/check-course-completion/:courseID", authorization, authorizationWithIDCheck, userProgressController.CheckCourseCompletion)
	router.POST("/mark-lessons-learned/:chapterID", authorization, userProgressController.MarkLessonsAsLearned)
	router.GET("/get-test-by-chapter/:chapter_id", authorization, userProgressController.GetTestByChapter) 
	router.POST("/submit-test-answers/:test_id", authorization, userProgressController.SubmitTestAnswers)
	router.GET("/get-test-result/:test_id", authorization, userProgressController.GetTestResult)
}