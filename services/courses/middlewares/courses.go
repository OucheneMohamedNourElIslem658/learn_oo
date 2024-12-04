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

		var count int64
		err := database.Model(models.Course{}).Where("id = ? and author_id = ?", courseID, authorID).Count(&count).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		if count == 0 {
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

func (cm *ChaptersMiddlewares) CheckChapterExistance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorID := ctx.GetString("author_id")
		courseID := ctx.Param("course_id")
		chapterID := ctx.Param("chapter_id")

		database := cm.database

		var count int64
		err := database.Model(models.Chapter{}).
		Select("courses.id, courses.author_id, chapters.id").
		Joins("JOIN chapters ON chapters.id = lessons.chapter_id").
		Joins("JOIN courses ON courses.id = chapters.course_id").
		Where("authors.id = ? AND courses.id = ? AND chapters.id", authorID, courseID, chapterID).
		Count(&count).
		Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		if count == 0 {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "chapter not found",
			})
			return
		}

		chapterIDInt, _ := strconv.Atoi(chapterID)

		ctx.Set("chapter_id", chapterIDInt)
		ctx.Next()
	}
}