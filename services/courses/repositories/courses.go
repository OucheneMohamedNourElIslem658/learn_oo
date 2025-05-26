package repositories

import (
	"fmt"
	"math"
	"mime/multipart"
	"net/http"
	"strings"

	gorm "gorm.io/gorm"

	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	filestorage "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/file_storage"
	models "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	payment "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/payment"
	utils "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
)

type CoursesRepository struct {
	database    *gorm.DB
	filestorage *filestorage.FileStorage
	payment     *payment.Payment
}

func NewCoursesRepository() *CoursesRepository {
	return &CoursesRepository{
		database:    database.Instance,
		filestorage: filestorage.NewFileStorage(),
		payment:     payment.NewPayment(),
	}
}

type CreatedCourseDTO struct {
	Title       string                `form:"title" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Price       float64               `form:"price" binding:"price"`
	Language    models.Languages      `form:"language" binding:"required,oneof='ar' 'fr' 'en'"`
	Level       models.Level          `form:"level" binding:"required,oneof='bigener' 'medium' 'advanced'"`
	Duration    float64               `form:"duration" binding:"omitempty,min=5"`
	Video       *multipart.FileHeader `form:"video,omitempty" binding:"required"`
	Image       *multipart.FileHeader `form:"image,omitempty" binding:"required"`
}

func (cr *CoursesRepository) CreateCourse(authorID string, course CreatedCourseDTO) (apiError *utils.APIError) {
	// Upload Image And Video:

	filestorage := cr.filestorage

	image, _ := course.Image.Open()
	defer image.Close()

	video, _ := course.Video.Open()
	defer video.Close()

	imageUploadResult, err := filestorage.UploadFile(image, fmt.Sprintf("/learn_oo/authors/%v/courses/images", authorID))
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	videoUploadResult, err := filestorage.UploadFile(video, fmt.Sprintf("/learn_oo/authors/%v/courses/videos", authorID))
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	courseToCreate := models.Course{
		AuthorID:    authorID,
		Title:       course.Title,
		Description: course.Description,
		Price:       course.Price,
		Duration:    uint(course.Duration),
		Language:    course.Language,
		Level:       course.Level,
		Image: &models.File{
			URL:          imageUploadResult.Url,
			ThumbnailURL: &imageUploadResult.ThumbnailUrl,
			Height:       imageUploadResult.Height,
			Width:        imageUploadResult.Width,
		},
		Video: &models.File{
			URL:          videoUploadResult.Url,
			ThumbnailURL: &videoUploadResult.ThumbnailUrl,
			Height:       videoUploadResult.Height,
			Width:        videoUploadResult.Width,
		},
	}

	if course.Price >= 50 {
		payment := cr.payment
		product, err := payment.CreateProduct(courseToCreate)
		if err != nil {
			return &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
		courseToCreate.PaymentPriceID = &product.PriceID
		courseToCreate.PaymentProductID = &product.ID
	}

	// Create Course:
	database := cr.database

	err = database.Create(&courseToCreate).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (cr *CoursesRepository) GetCourse(ID, userID, authorID, appendWith string) (course *models.Course, apiError *utils.APIError) {
	database := cr.database

	query := database.Model(&models.Course{}).Where("id = ?", ID)

	validExtentions := utils.GetValidExtentions(
		appendWith,
		"author",
		"image",
		"video",
		"requirements",
		"objectives",
		"categories",
		"chapters",
		"learners",
	)

	for _, extention := range validExtentions {
		switch extention {
		case "Chapters":
			query = query.Preload(extention, func(db *gorm.DB) *gorm.DB {
				return db.Preload("Test", func(db *gorm.DB) *gorm.DB {
					return db.Select("tests.*, COUNT(questions.id) AS questions_count, "+
						"CASE WHEN test_results.test_id IS NOT NULL AND test_results.has_succeed = TRUE THEN TRUE ELSE FALSE END AS has_succeed").
						Joins("LEFT JOIN test_results ON test_results.test_id = tests.id AND test_results.learner_id = ?", userID).
						Joins("LEFT JOIN questions ON questions.test_id = tests.id").
						Group("tests.id, test_results.test_id, test_results.has_succeed")
				}).Preload("Lessons", func(db *gorm.DB) *gorm.DB {
					return db.Select("lessons.id, lessons.title, lessons.description, lessons.chapter_id, "+
						"CASE WHEN files.id IS NOT NULL THEN TRUE ELSE FALSE END AS is_video, "+
						"CASE WHEN lesson_learners.lesson_id IS NOT NULL AND lesson_learners.learned = TRUE THEN TRUE ELSE FALSE END AS learned").
						Joins("LEFT JOIN files ON lessons.id = files.lesson_id").
						Joins("LEFT JOIN lesson_learners ON lesson_learners.lesson_id = lessons.id AND lesson_learners.learner_id = ?", userID) // Assuming `userID` is available
				})
			})
		case "Author":
			query.Preload("Author.User")
		default:
			query.Preload(extention)
		}
	}

	if authorID != "" {
		query.Where("author_id = ?", authorID)
	} else {
		query.Where("is_completed = ?", true)
	}

	query.Select("courses.*, COALESCE(AVG(course_learners.rate), 0) AS rate").
		Joins("LEFT JOIN course_learners ON course_learners.course_id = courses.id").
		Group("courses.id")

	var existingCourse models.Course
	err := query.First(&existingCourse).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "course not found",
			}
		}
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return &existingCourse, nil
}

type CourseSearchDTO struct {
	Title           string           `form:"title"`
	FreeOrPaid      string           `form:"free_or_paid" binding:"omitempty,oneof='free' 'paid'"`
	Language        models.Languages `form:"language" binding:"omitempty,oneof='ar' 'fr' 'en'"`
	Level           models.Level     `form:"level" binding:"omitempty,oneof='bigener' 'medium' 'advanced'"`
	MinDuration     float64          `form:"min_duration" binding:"omitempty,min=5"`
	MaxDuration     float64          `form:"max_duration" binding:"omitempty,min=5"`
	PageSize        uint             `form:"page_size,default=10" binding:"min=1"`
	Page            uint             `form:"page,default=1" binding:"min=1"`
	AppendWith      string           `form:"append_with"`
	CategoriesNames string           `form:"categories_names"`
}

func (cr *CoursesRepository) GetCourses(filters CourseSearchDTO) (courses []models.Course, currentPage, count, maxPages *uint, apiError *utils.APIError) {
	database := cr.database

	query := database.Model(&models.Course{}).Where("is_completed = true")

	title := filters.Title
	language := filters.Language
	level := filters.Level
	minDuration := filters.MinDuration
	maxDuration := filters.MaxDuration
	appendWith := filters.AppendWith
	freePaid := filters.FreeOrPaid
	pageSize := filters.PageSize
	page := filters.Page

	var categoriesNames []string
	if len(filters.CategoriesNames) > 0 {
		categoriesNames = strings.Split(filters.CategoriesNames, ",")
	}

	validExtentions := utils.GetValidExtentions(
		appendWith,
		"author",
		"image",
		"video",
		"categories",
	)

	for _, extention := range validExtentions {
		if extention == "Author" {
			query.Preload("Author.User")
		} else {
			query.Preload(extention)
		}
	}

	if title != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(title)+"%")
	}

	if freePaid == "free" {
		query = query.Where("price = ?", 0)
	} else if freePaid == "paid" {
		query = query.Where("price <> ?", 0)
	}

	if language != "" {
		query = query.Where("language = ?", language)
	}

	if level != "" {
		query = query.Where("level = ?", level)
	}

	if minDuration > 0 {
		query = query.Where("duration >= ?", minDuration)
	}

	if maxDuration > 0 {
		query = query.Where("duration <= ?", maxDuration)
	}

	if len(categoriesNames) > 0 {
		query = query.Joins("JOIN course_categories ON course_categories.course_id = courses.id").
			Joins("JOIN categories ON course_categories.category_id = categories.id").
			Where("categories.name IN (?)", categoriesNames)
	}

	query.Select(`courses.*, 
		COALESCE(AVG(course_learners.rate), 0) AS rate, 
		SUM(CASE WHEN course_learners.rate IS NOT NULL THEN 1 ELSE 0 END) AS raters_count`).
		Joins("LEFT JOIN course_learners ON course_learners.course_id = courses.id").
		Group("courses.id").
		Order("rate DESC, price DESC, created_at DESC, duration DESC")

	var totalRecords int64
	database.Model(&models.User{}).Count(&totalRecords)
	totalPages := uint(math.Ceil(float64(totalRecords) / float64(pageSize)))

	offset := (page - 1) * pageSize
	query.Limit(int(pageSize)).Offset(int(offset))

	var coursesList []models.Course
	err := query.Find(&coursesList).Error
	if err != nil {
		return nil, nil, nil, nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	coursesListLenght := uint(len(coursesList))

	return coursesList, &page, &coursesListLenght, &totalPages, nil
}

type UpdateCourseDTO struct {
	Title           string           `json:"title"`
	Description     string           `json:"description"`
	Price           *float64         `json:"price" binding:"omitempty,price"`
	Language        models.Languages `json:"language" binding:"omitempty,oneof='ar' 'fr' 'en'"`
	Level           models.Level     `json:"level" binding:"omitempty,oneof='bigener' 'medium' 'advanced'"`
	Duration        float64          `form:"duration" binding:"omitempty,min=5"`
	IsCompleted     *bool            `form:"is_completed"`
	CategoriesNames []string         `json:"categories_names"`
}

func (cr *CoursesRepository) UpdateCourse(ID, authorID string, course UpdateCourseDTO) (apiError *utils.APIError) {
	database := cr.database

	var existingCourse models.Course
	err := database.Where("id = ? and author_id = ?", ID, authorID).Preload("Image").Preload("Video").First(&existingCourse).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "course not found",
			}
		}
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if course.Title != "" {
		existingCourse.Title = course.Title
	}

	if course.Description != "" {
		existingCourse.Description = course.Description
	}

	if course.Language != "" {
		existingCourse.Language = course.Language
	}

	if course.Level != "" {
		existingCourse.Level = course.Level
	}

	if course.Duration > 0 {
		existingCourse.Duration = uint(course.Duration)
	}

	if course.IsCompleted != nil {
		existingCourse.IsCompleted = *course.IsCompleted
	}

	if course.Price != nil {
		existingCourse.Price = *course.Price
		if *course.Price >= 50 {
			payment := cr.payment
			product, err := payment.CreateProduct(existingCourse)
			if err != nil {
				return &utils.APIError{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
				}
			}
			existingCourse.PaymentPriceID = &product.PriceID
			existingCourse.PaymentProductID = &product.ID
		}
	}

	if len(course.CategoriesNames) != 0 {
		categoriesNames := course.CategoriesNames
		var categories []models.Category
		err = database.Where("name IN (?)", categoriesNames).Find(&categories).Error
		if err != nil {
			return &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}

		err = database.Model(&existingCourse).Association("Categories").Replace(categories)
		if err != nil {
			return &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
	}

	err = database.Save(&existingCourse).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (cr *CoursesRepository) UpdateCourseImage(ID uint, authorID string, image multipart.File) (apiError *utils.APIError) {
	database := cr.database
	filestorage := cr.filestorage

	var existingImage models.File
	err := database.Where("image_course_id = ?", ID).First(&existingImage).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if err == nil {
		if err := database.Where("id = ?", existingImage.ID).Unscoped().Delete(&existingImage).Error; err != nil {
			return &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
		if existingImage.ImageKitID != nil {
			if err := filestorage.DeleteFile(*existingImage.ImageKitID); err != nil {
				return &utils.APIError{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
				}
			}
		}
	}

	path := fmt.Sprintf("/learn_oo/authors/%v/courses/images", authorID)
	uploadData, err := filestorage.UploadFile(image, path)
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	newImage := models.File{
		URL:           uploadData.Url,
		ImageKitID:    &uploadData.FileId,
		ThumbnailURL:  &uploadData.ThumbnailUrl,
		ImageCourseID: &ID,
		Height:        uploadData.Height,
		Width:         uploadData.Width,
	}
	if err := database.Create(&newImage).Error; err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (cr *CoursesRepository) UpdateCourseVideo(ID uint, authorID string, video multipart.File) (apiError *utils.APIError) {
	database := cr.database
	filestorage := cr.filestorage

	var existingVideo models.File
	err := database.Where("video_course_id = ?", ID).First(&existingVideo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if err == nil {
		if err := database.Where("id = ?", existingVideo.ID).Unscoped().Delete(&existingVideo).Error; err != nil {
			return &utils.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
		if existingVideo.ImageKitID != nil {
			if err := filestorage.DeleteFile(*existingVideo.ImageKitID); err != nil {
				return &utils.APIError{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
				}
			}
		}
	}

	path := fmt.Sprintf("/learn_oo/authors/%v/courses/videos", authorID)
	uploadData, err := filestorage.UploadFile(video, path)
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	newVideo := models.File{
		URL:           uploadData.Url,
		ImageKitID:    &uploadData.FileId,
		ThumbnailURL:  &uploadData.ThumbnailUrl,
		VideoCourseID: &ID,
		Height:        uploadData.Height,
		Width:         uploadData.Width,
	}
	if err := database.Create(&newVideo).Error; err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (cr *CoursesRepository) DeleteCourse(ID, authorID string) (apiError *utils.APIError) {
	database := cr.database

	var existingCourse models.Course
	err := database.Where("id = ? and author_id = ?", ID, authorID).First(&existingCourse).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "course not found",
			}
		}
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	err = database.Unscoped().Delete(&existingCourse).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (cr *CoursesRepository) GetCategories() ([]models.Category, *utils.APIError) {
	database := cr.database

	var categories []models.Category

	err := database.Find(&categories).Error
	if err != nil {
		return nil, &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return categories, nil
}

type CreatedCategoryDTO struct {
	Name string `json:"name" binding:"required"`
}

func (cr *CoursesRepository) CreateCategory(category CreatedCategoryDTO) (apiError *utils.APIError) {
	database := cr.database

	var existingCategory models.Category

	err := database.Where("name = ?", category.Name).First(&existingCategory).Error
	if err == nil {
		return &utils.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    "categoriy name already exists",
		}
	}

	newCategory := models.Category{Name: category.Name}
	err = database.Create(&newCategory).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (cr *CoursesRepository) DeleteCategory(ID string) *utils.APIError {
	database := cr.database

	var category models.Category

	err := database.Where("id = ?", ID).First(&category).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &utils.APIError{
				StatusCode: http.StatusNotFound,
				Message:    "category not found",
			}
		}
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	err = database.Delete(&category).Error
	if err != nil {
		return &utils.APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}
