package repositorycomment

import "gorm.io/gorm"

type CommentModel struct {
	gorm.Model
	PostingID uint
	UserID    uint
	Comment   string
}
