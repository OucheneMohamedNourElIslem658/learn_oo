package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/repositories"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
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

		if authorization != "" {
			idToken := authorization[len("Bearer "):]

			if idToken != "" {
				claims, isValid, err := utils.VerifyIDToken(idToken)

				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": "invalid id token",
					})
					return
				}

				if !isValid {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"error": "id token expired",
					})
					return
				}

				ctx.Set("id", claims.ID)
				// ctx.Set("email_verified", claims.EmailVerified)
				if claims.AuthorID != nil {
					ctx.Set("author_id", *claims.AuthorID)
				}
			}
		}
		ctx.Next()
	}
}

func (am *AuthorizationMiddlewares) AuthorizationWithIDCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorID := ctx.GetString("id")

		if authorID == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "requester is not a user",
			})
			return
		}

		ctx.Next()
	}
}

// func (am *AuthorizationMiddlewares) AuthorizationWithEmailVerification() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		emailVerified := ctx.GetBool("email_verified")

// 		if !emailVerified {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error": "email is not verified",
// 			})
// 			return
// 		}

// 		ctx.Next()
// 	}
// }

func (am *AuthorizationMiddlewares) AuthorizationWithAuthorCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorID := ctx.GetString("author_id")
		if authorID == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "user is not author",
			})
			return
		}

		ctx.Next()
	}
}

func (am *AuthorizationMiddlewares) AuthorizationWithEmailCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idToken := ctx.Param("id_token")
		email, isValid, err := utils.VerifyIDTokenFromEmail(idToken)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if !isValid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token expired",
			})
			return
		}

		ctx.Set("email", *email)
		ctx.Next()
	}
}
