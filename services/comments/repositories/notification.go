package repositories

import (
	"gorm.io/gorm"
	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	models "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
)

// NotificationRepository handles the database operations related to notifications.
type NotificationRepository struct {
	database *gorm.DB
}

// NewNotificationRepository creates a new instance of NotificationRepository.
func NewNotificationRepository() *NotificationRepository {
	return &NotificationRepository{
		database: database.Instance,
	}
}

// CreateMany inserts multiple notifications into the database.
func (repository *NotificationRepository) CreateMany(notifications []models.Notification) error {
	if err := repository.database.Create(&notifications).Error; err != nil {
		return err
	}
	return nil
}

// GetNotificationsByUserID retrieves all notifications for a user.
func (repository *NotificationRepository) GetNotificationsByUserID(userID string) ([]models.Notification, error) {
	var notifications []models.Notification
	err := repository.database.Unscoped().Where("user_id = ?", userID).Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

// func (repository *NotificationRepository) GetNotificationByID(notificationID uint) (*models.Notification, error) {
// 	var notification models.Notification
// 	if err := repository.database.First(&notification, notificationID).Error; err != nil {
// 		return nil, err
// 	}
// 	return &notification, nil
// }


