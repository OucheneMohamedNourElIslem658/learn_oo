package repositories

import (
	"gorm.io/gorm"
    database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
    models "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"

)

type CommentsRepository struct {
	database *gorm.DB
}

func NewCommentsRepository(db *gorm.DB) *CommentsRepository {
	return &CommentsRepository{database: database.Instance}
}

type CreateCommentRequest struct {
    Content   string `json:"content" binding:"required"`
    RepliedTo *uint  `json:"replied_to"`
    UserID    string   `json:"-"`
    LessonID  uint   `json:"-"`
}

func (repository *CommentsRepository) CreateComment(request CreateCommentRequest) (*models.Comment, error) {
    comment := &models.Comment{
        Content:  request.Content,
        LessonID: request.LessonID,
        UserID:   request.UserID,
    }
    if request.RepliedTo != nil {
        comment.RepliedTo = request.RepliedTo
    }
    if err := repository.database.Create(comment).Error; err != nil {
        return nil, err
    }
    return comment, nil
}

func (r *CommentsRepository) GetByID(id uint, comment *models.Comment) error {
    return r.database.First(comment, id).Error
}

// GetByLessonID retrieves all comments for a given lesson ID
func (r *CommentsRepository) GetByLessonID(lessonID uint) ([]models.Comment, error) {
    var comments []models.Comment
    err := r.database.Where("lesson_id = ?", lessonID).Find(&comments).Error
    return comments, err
}

// GetByUserID retrieves all comments made by a user
func (r *CommentsRepository) GetByUserID(userID string) ([]models.Comment, error) {
    var comments []models.Comment
    err := r.database.Where("user_id = ?", userID).Find(&comments).Error
    return comments, err
}

// Delete  a comment
func (r *CommentsRepository) Delete(id uint, userID string) error {
    // Optionally add checks to ensure the user is the owner of the comment
    return r.database.Delete(&models.Comment{}, "id = ? AND user_id = ?", id, userID).Error
}