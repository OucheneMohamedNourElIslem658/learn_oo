package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/comments/controllers"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/middlewares"
)

type CommentsRouter struct {
	Router                   *gin.RouterGroup
	commentsController       *controllers.CommentsController
	notificationsController  *controllers.NotificationsController
	authMiddlewares          *middlewares.AuthorizationMiddlewares
}

func NewCommentsRouter(router *gin.RouterGroup) *CommentsRouter {
	// Initialize the controllers directly, which will internally handle the DB connection
	return &CommentsRouter{
		Router:                  router,
		commentsController:      controllers.NewCommentsController(),
		notificationsController: controllers.NewNotificationsController(),
		authMiddlewares:         middlewares.NewAuthorizationMiddlewares(),
	}
}

func (cr *CommentsRouter) RegisterRoutes() {
	router := cr.Router
	commentsController := cr.commentsController
	notificationsController := cr.notificationsController
	authMiddlewares := cr.authMiddlewares

	authorization := authMiddlewares.Authorization()

	// Comment routes
	commentsRouter := router.Group("/comments")
	commentsRouter.POST("/:lesson_id", authorization, commentsController.Create)
	commentsRouter.GET("/:id",authorization, commentsController.GetByID)  // Get comment by ID
    commentsRouter.GET("/lesson/:lesson_id", authorization,commentsController.GetByLessonID) // Get comments by lesson ID
    commentsRouter.GET("/user", authorization, commentsController.GetByUserID)  // Get comments by user ID
    commentsRouter.DELETE("/:id", authorization, commentsController.Delete) // Delete comment

	// Notification routes
	notificationsRouter := router.Group("/notifications")
	notificationsRouter.GET("/all_notification", notificationsController.GetNotificationsByUserID) // Get notifications by user ID
}