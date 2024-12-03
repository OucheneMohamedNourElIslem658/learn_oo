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
	authorizationWithEmailVerification := authMiddlewares.AuthorizationWithEmailVerification()
	authorizationWithAuthorCheck := authMiddlewares.AuthorizationWithAuthorCheck()

	profileRouter := router.Group("/profile")
	profileRouter.GET("/", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, usersController.GetUser)
	profileRouter.PUT("/image", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, usersController.UpdateUserImage)
	profileRouter.PUT("/", authorization, authorizationWithEmailVerification, usersController.UpdateUser)

	authorsRouter := router.Group("/authors")
	authorsRouter.PUT("/upgrade", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, usersController.UpgradeToAuthor)
	authorsRouter.DELETE("/downgrade", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, authorizationWithAuthorCheck, usersController.DowngradeFromAuthor)

	authorProfileRouter := authorsRouter.Group("/profile")
	authorProfileRouter.GET("/", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, authorizationWithAuthorCheck, usersController.GetAuthor)
	authorProfileRouter.PUT("/", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, authorizationWithAuthorCheck, usersController.UpdateAuthor)

	authorAccomplishments := authorProfileRouter.Group("/accomplishments")
	authorAccomplishments.POST("/", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, authorizationWithAuthorCheck, usersController.AddAuthorAccomplishments)
	authorAccomplishments.DELETE("/:file_id", authorization, authorizationWithIDCheck, authorizationWithEmailVerification, authorizationWithAuthorCheck, usersController.DeleteAuthorAccomplishment)
}
