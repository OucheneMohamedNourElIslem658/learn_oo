package middlewares

import (
	// "fmt"
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
			ctx.Abort()
		}

		idToken := authorization[len("Bearer "):]
		claims, isValid, err := utils.VerifyIDToken(idToken)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			ctx.Abort()
		}

		if !isValid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "id token expired",
			})
			ctx.Abort()
		}

		ctx.Set("id", claims.ID)
		ctx.Set("author_id", *claims.AuthorID)
		ctx.Set("email_verified", claims.EmailVerified)
		ctx.Next()
	}
}

func (am *AuthorizationMiddlewares) AuthorizationWithEmailVerification() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		emailVerified := ctx.GetBool("email_verified")

		if !emailVerified {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "email is not verified",
			})
			ctx.Abort()
		}

		ctx.Next()
	}
}

func (am *AuthorizationMiddlewares) AuthorizationWithAuthorCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorID := ctx.GetString("author_id")

		if authorID == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "user is not author",
			})
			ctx.Abort()
		}

		ctx.Next()
	}
}

func (am *AuthorizationMiddlewares) AuthorizationWithEmailCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idToken := ctx.Param("id_token")
		email, isValid, err := utils.VerifyIDTokenFromEmail(idToken)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
		}

		if !isValid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "token expired",
			})
			ctx.Abort()
		}

		ctx.Set("email", *email)
		ctx.Next()
	}
}
