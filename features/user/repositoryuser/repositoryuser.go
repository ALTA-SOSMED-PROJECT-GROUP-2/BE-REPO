package repositoryuser

import (
	"BE-REPO/features/comment/repositorycomment"
	"BE-REPO/features/posting/repositoryposting"
	"BE-REPO/features/user"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email         string
	FullName      string
	UserName      string
	PhoneNumber   string
	Image         string
	Password      string
	PostingModel  []repositoryposting.PostingModel `gorm:"foreignKey:UserID"`
	CommentModels []repositorycomment.CommentModel `gorm:"foreignKey:UserID"`
}

type userQuery struct {
	DB *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &userQuery{
		DB: db,
	}
}

func (uq *userQuery) Insert(newUser user.User) (user.User, error) {
	var inputDB = new(UserModel)
	inputDB.Email = newUser.Email
	inputDB.PhoneNumber = newUser.PhoneNumber
	inputDB.FullName = newUser.FullName
	inputDB.UserName = newUser.UserName
	inputDB.Password = newUser.Password

	if err := uq.DB.Create(&inputDB).Error; err != nil {
		return user.User{}, err
	}

	newUser.ID = inputDB.ID

	return newUser, nil
}

func (uq *userQuery) Login(email string, password string) (user.User, error) {
	var userData = new(UserModel)

	if err := uq.DB.Where("email = ? AND password = ?", email, password).First(userData).Error; err != nil {
		return user.User{}, err
	}

	var result = new(user.User)
	result.ID = userData.ID
	result.Email = userData.Email
	result.Password = userData.Password

	return *result, nil
}

func (uq *userQuery) GetUserByEmail(email string) (user.User, error) {
	var userData = new(UserModel)

	if err := uq.DB.Where("email = ?", email).First(userData).Error; err != nil {
		return user.User{}, err
	}

	var result = new(user.User)
	result.ID = userData.ID
	result.Email = userData.Email
	result.Password = userData.Password

	return *result, nil
}
