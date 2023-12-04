package repositoryuser

import (
	"BE-REPO/features/comment/repositorycomment"
	"BE-REPO/features/posting/repositoryposting"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email         string
	FullName      string
	PhoneNumber   string
	Password      string
	PostingModel  []repositoryposting.PostingModel `gorm:"foreignKey:UserID"`
	CommentModels []repositorycomment.CommentModel `gorm:"foreignKey:UserID"`
}
