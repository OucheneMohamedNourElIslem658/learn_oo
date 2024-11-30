package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
)

type ChaptersMiddlewares struct {
	database *gorm.DB
}

func NewChaptersMiddlewares() *ChaptersMiddlewares {
	return &ChaptersMiddlewares{
		database: database.Instance,
	}
}

func (cm *ChaptersMiddlewares) CheckCourseExistance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorID := ctx.GetString("author_id")
		courseID := ctx.Param("course_id")

		database := cm.database

		var exists bool
		err := database.Model(models.Course{}).Where("course_id = ? and author_id = ?", courseID, authorID).Scan(&exists).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		if !exists {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "course not found",
			})
			return
		}

		courseIDInt, _ := strconv.Atoi(courseID)

		ctx.Set("course_id", uint(courseIDInt))
		ctx.Next()
	}
}