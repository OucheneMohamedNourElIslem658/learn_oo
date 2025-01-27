package controllers

import (
	"log"
	"net/http"
	"strconv"
	models "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	repositories "github.com/OucheneMohamedNourElIslem658/learn_oo/services/comments/repositories"
	"github.com/gin-gonic/gin"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database" // Import the database package
)

type CommentsController struct {
	commentsRepository     *repositories.CommentsRepository
	notificationRepository *repositories.NotificationRepository
}

func NewCommentsController() *CommentsController {
	return &CommentsController{
		commentsRepository:     repositories.NewCommentsRepository(database.Instance), 
		notificationRepository: repositories.NewNotificationRepository(),
	}
}

func (c *CommentsController) Create(ctx *gin.Context) {
	var body repositories.CreateCommentRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lessonIDParam := ctx.Param("lesson_id")
	lessonID, err := strconv.Atoi(lessonIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid lesson ID"})
		return
	}
	body.LessonID = uint(lessonID)

	userIDStr := ctx.GetString("id")

	
	body.UserID = userIDStr

	comment, err := c.commentsRepository.CreateComment(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var lesson models.Lesson
	if err := database.Instance.Preload("Learners").First(&lesson, comment.LessonID).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var notifications []models.Notification
	commentUserIDStr :=comment.UserID

	for _, learner := range lesson.Learners {
		if learner.ID != commentUserIDStr {
			notifications = append(notifications, models.Notification{
				Title:       "New Comment",
				Description: "A new comment was added to lesson: " + lesson.Title,
				CommentID:   &comment.ID,
				UserID:      learner.ID,
			})
		}
	}

	if len(notifications) > 0 {
		if err := c.notificationRepository.CreateMany(notifications); err != nil {
			log.Println("Failed to create notifications:", err)
		}
	}

	ctx.JSON(http.StatusCreated, comment)
}

func (c *CommentsController) GetByID(ctx *gin.Context) {
    commentIDParam := ctx.Param("id")
    commentID, err := strconv.Atoi(commentIDParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment ID"})
        return
    }

    var comment models.Comment
    if err := c.commentsRepository.GetByID(uint(commentID), &comment); err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
        return
    }

    ctx.JSON(http.StatusOK, comment)
}



func (c *CommentsController) GetByLessonID(ctx *gin.Context) {
    lessonIDParam := ctx.Param("lesson_id")
    lessonID, err := strconv.Atoi(lessonIDParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid lesson ID"})
        return
    }

    comments, err := c.commentsRepository.GetByLessonID(uint(lessonID))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, comments)
}


func (c *CommentsController) GetByUserID(ctx *gin.Context) {
    // Retrieve userID from the context (set by the middleware)
    userIDStr := ctx.GetString("id") // Use GetString to retrieve the user ID as a string
    

    // Convert the userID string to uint
    // userID, err := strconv.ParseUint(userIDStr, 10, 32)
    // if err != nil {
    //     ctx.JSON(http.StatusBadRequest, gin.H{"error": err.rror()})
    //     return
    // }

    // Retrieve comments by userID
    comments, err := c.commentsRepository.GetByUserID(userIDStr)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Return the comments
    ctx.JSON(http.StatusOK, comments)
}


func (c *CommentsController) Delete(ctx *gin.Context) {
    commentIDParam := ctx.Param("id")
    commentID, err := strconv.Atoi(commentIDParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment ID"})
        return
    }

    userID := ctx.GetString("id")

    err = c.commentsRepository.Delete(uint(commentID), userID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "comment deleted successfully"})
}