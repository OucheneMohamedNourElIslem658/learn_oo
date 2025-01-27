package controllers

import (
	"net/http"
	

	"github.com/gin-gonic/gin"
	repositories "github.com/OucheneMohamedNourElIslem658/learn_oo/services/comments/repositories"
)

type NotificationsController struct {
	notificationRepository *repositories.NotificationRepository
}

func NewNotificationsController() *NotificationsController {
	return &NotificationsController{
		notificationRepository: repositories.NewNotificationRepository(),
	}
}

func (nc *NotificationsController) GetNotificationsByUserID(ctx *gin.Context) {
	userID:= ctx.GetString("id")

	// userIDStr, ok := userID.(string)
	// if !ok {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type in context"})
	// 	return
	// }

	// // Convert the user ID to uint
	// userIDUint, err := strconv.ParseUint(userIDStr, 10, 32)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
	// 	return
	// }

	notifications, err := nc.notificationRepository.GetNotificationsByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notifications)
}