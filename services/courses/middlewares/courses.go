package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
)

type CoursesMiddlewares struct {
	database *gorm.DB
}

func NewCoursesMiddlewares() *CoursesMiddlewares {
	return &CoursesMiddlewares{
		database: database.Instance,
	}
}

func (cm *CoursesMiddlewares) CheckCourseExistance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorID := ctx.GetString("author_id")
		courseID := ctx.Param("course_id")

		database := cm.database

		var count int64
		err := database.Model(models.Course{}).Where("id = ? and author_id = ?", courseID, authorID).Count(&count).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if count == 0 {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "course not found",
			})
			return
		}

		courseIDInt, _ := strconv.Atoi(courseID)

		ctx.Set("course_id", uint(courseIDInt))
		ctx.Next()
	}
}

func (cm *CoursesMiddlewares) CheckChapterExistance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorID := ctx.GetString("author_id")
		courseID := ctx.Param("course_id")
		chapterID := ctx.Param("chapter_id")

		database := cm.database

		var count int64
		err := database.Model(&models.Chapter{}).
			Select("courses.id, courses.author_id, chapters.id").
			Joins("JOIN courses ON courses.id = chapters.course_id").
			Where("courses.author_id = ? AND courses.id = ? AND chapters.id = ?", authorID, courseID, chapterID).
			Count(&count).
			Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if count == 0 {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "chapter not found",
			})
			return
		}

		chapterIDInt, _ := strconv.Atoi(chapterID)

		ctx.Set("chapter_id", chapterIDInt)
		ctx.Next()
	}
}