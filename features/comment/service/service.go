package service

import (
	"BE-REPO/features/comment/model"
	"BE-REPO/features/comment/repository"
)

type CommentService struct {
	cr *repository.CommentRepository
}

func NewCommentService(cr *repository.CommentRepository) *CommentService {
	return &CommentService{cr: cr}
}

func (cs *CommentService) Create(comment *model.Comment) error {
	return cs.cr.Create(comment)
}

func (cs *CommentService) Get(id uint) (*model.Comment, error) {
	return cs.cr.Get(id)
}

func (cs *CommentService) GetAll(postingID uint) ([]model.Comment, error) {
	return cs.cr.GetAll(postingID)
}

func (cs *CommentService) Delete(id uint) error {
	return cs.cr.Delete(id)
}
