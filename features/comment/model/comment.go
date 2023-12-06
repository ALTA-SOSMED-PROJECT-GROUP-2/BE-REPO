package model

type Comment struct {
	ID        uint   `gorm:"primary_key"`
	PostingID uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	Content   string `gorm:"not null"`
}
