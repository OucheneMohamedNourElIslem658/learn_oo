package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/repositories"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/utils"
)

type AuthorizationMiddlewares struct {
	authRepository *repositories.AuthRepository
}

func NewAuthorizationMiddlewares() *AuthorizationMiddlewares {
	return &AuthorizationMiddlewares{
		authRepository: repositories.NewAuthRepository(),
	}
}

func (am *AuthorizationMiddlewares) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")

		if authorization == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "undefined authorization",
			})
			return
		}
	
		idToken := authorization[len("Bearer "):]
		claims, isValid, err := utils.VerifyIDToken(idToken)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}

		if !isValid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "id token expired",
			})
			return
		}

		ctx.Set("id", claims.ID)
		ctx.Set("author_id", claims.AuthorID)
		ctx.Next()
	}
}

func (am *AuthorizationMiddlewares) AuthorizationWithEmailVerification() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorID, _ := ctx.Get("author_id")

		if authorID.(*string) == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "user is not author",
			})
			return
		}

		ctx.Next()
	}
}

func (am *AuthorizationMiddlewares) AuthorizationWithAuthorCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorID, _ := ctx.Get("author_id")

		if authorID.(*string) == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "user is not author",
			})
			return
		}

		ctx.Next()
	}
}
