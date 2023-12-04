package repositoryposting

import "gorm.io/gorm"

type PostingModel struct {
	gorm.Model
	Title       string
	Description string
	Image       string
	UserID      uint
}
