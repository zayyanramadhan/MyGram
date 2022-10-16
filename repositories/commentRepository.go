package repositories

import (
	"final_MyGram/models"
	"log"

	"gorm.io/gorm"
)

type CommentRepo interface {
	CreateCommentRepository(authId int, comment *models.Comment) (*models.Comment, error)
	GetAllCommentRepository() (*[]models.CommentResponse, error)
	UpdateCommentRepository(authId int, comment *models.Comment) (*models.Comment, error)
	DeleteCommentRepository(commentId, authId int) error
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepo {
	return &commentRepo{db}
}

func (r *commentRepo) CreateCommentRepository(authId int, comment *models.Comment) (*models.Comment, error) {
	comment.UserId = authId
	err := r.db.Create(&comment)
	if err.Error != nil {
		return nil, err.Error
	}
	return comment, nil
}

func (r *commentRepo) GetAllCommentRepository() (*[]models.CommentResponse, error) {
	comments := []models.CommentResponse{}
	if err := r.db.Joins("User").Joins("Photo").Model(&models.Comment{}).Find(&comments).Error; err != nil {
		log.Printf("Error get data with err: %s", err)
		return nil, err
	}
	return &comments, nil
}

func (r *commentRepo) UpdateCommentRepository(authId int, comment *models.Comment) (*models.Comment, error) {
	commentCheckAvail := models.Comment{}
	err := r.db.Where("user_id = ?", authId).
		Where("id = ?", comment.Id).
		First(&commentCheckAvail)
	if err.Error != nil {
		log.Printf("Error get data comment detail with err: %s", err.Error)
		return nil, err.Error
	}
	err = r.db.Model(&commentCheckAvail).Updates(comment)
	if err.Error != nil {
		log.Printf("Error update data comment id with err: %s", err.Error)
		return nil, err.Error
	}
	return &commentCheckAvail, nil
}

func (r *commentRepo) DeleteCommentRepository(commentId, authId int) error {
	commentCheckAvail := models.Comment{}
	err := r.db.Where("user_id = ?", authId).
		Where("id = ?", commentId).
		First(&commentCheckAvail)
	if err.Error != nil {
		log.Printf("Error delete comment detail with err: %s", err.Error)
		return err.Error
	}
	r.db.Delete(commentCheckAvail, commentId)
	return nil
}
