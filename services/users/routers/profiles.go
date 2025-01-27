package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/controllers"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/middlewares"
)

type ProfilesRouter struct {
	Router          *gin.RouterGroup
	usersController *controllers.ProfilesController
	authMiddlewares *middlewares.AuthorizationMiddlewares
}

func NewProfilesRouter(router *gin.RouterGroup) *ProfilesRouter {
	return &ProfilesRouter{
		Router:          router,
		usersController: controllers.NewProfilesController(),
		authMiddlewares: middlewares.NewAuthorizationMiddlewares(),
	}
}

func (pr *ProfilesRouter) RegisterRoutes() {
	router := pr.Router
	usersController := pr.usersController

	authMiddlewares := pr.authMiddlewares
	authorization := authMiddlewares.Authorization()
	authorizationWithIDCheck := authMiddlewares.AuthorizationWithIDCheck()
	authorizationWithAuthorCheck := authMiddlewares.AuthorizationWithAuthorCheck()

	profileRouter := router.Group("/profile")
	profileRouter.GET("/", authorization, authorizationWithIDCheck, usersController.GetUser)
	profileRouter.PUT("/image", authorization, authorizationWithIDCheck, usersController.UpdateUserImage)
	profileRouter.PUT("/", authorization, authorizationWithIDCheck, usersController.UpdateUser)

	authorsRouter := router.Group("/authors")
	authorsRouter.PUT("/upgrade", authorization, authorizationWithIDCheck, usersController.UpgradeToAuthor)
	authorsRouter.DELETE("/downgrade", authorization, authorizationWithAuthorCheck, usersController.DowngradeFromAuthor)
	authorsRouter.GET("/:author_id", usersController.GetAuthor)

	authorProfileRouter := authorsRouter.Group("/profile")
	authorProfileRouter.GET("/", authorization, authorizationWithAuthorCheck, usersController.GetAuthorPorfile)
	authorProfileRouter.PUT("/", authorization, authorizationWithAuthorCheck, usersController.UpdateAuthor)

	authorAccomplishments := authorProfileRouter.Group("/accomplishments")
	authorAccomplishments.POST("/", authorization, authorizationWithAuthorCheck, usersController.AddAuthorAccomplishments)
	authorAccomplishments.DELETE("/:file_id", authorization, authorizationWithAuthorCheck, usersController.DeleteAuthorAccomplishment)
}