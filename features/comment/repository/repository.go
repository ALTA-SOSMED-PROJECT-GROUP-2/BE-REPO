package repository

import (
	"BE-REPO/features/comment/model"

	"github.com/jinzhu/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (cr *CommentRepository) Create(comment *model.Comment) error {
	return cr.db.Create(comment).Error
}

func (cr *CommentRepository) Get(id uint) (*model.Comment, error) {
	var comment model.Comment
	err := cr.db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (cr *CommentRepository) GetAll(postingID uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := cr.db.Where("posting_id = ?", postingID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (cr *CommentRepository) Delete(id uint) error {
	var comment model.Comment
	err := cr.db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return err
	}
	return cr.db.Delete(&comment).Error
}
