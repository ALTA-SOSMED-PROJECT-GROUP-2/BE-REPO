package serviceuser

import (
	"BE-REPO/features/user"
	"errors"
	"strings"
)

type userService struct {
	repo user.Repository
}

func New(r user.Repository) user.Service {
	return &userService{
		repo: r,
	}
}

func (us *userService) Register(newUser user.User) (user.User, error) {
	if newUser.Email == "" {
		return user.User{}, errors.New("password cannot be empty")
	}
	if newUser.FullName == "" {
		return user.User{}, errors.New("password cannot be empty")
	}
	if newUser.PhoneNumber == "" {
		return user.User{}, errors.New("password cannot be empty")
	}
	if newUser.UserName == "" {
		return user.User{}, errors.New("password cannot be empty")
	}
	if newUser.Password == "" {
		return user.User{}, errors.New("password cannot be empty")
	}

	result, err := us.repo.Insert(newUser)
	if err != nil {
		return user.User{}, errors.New("telah terjadi kesalahan terhadap sistem")
	}

	return result, nil
}

func (us *userService) Login(email string, password string) (user.User, error) {
	result, err := us.repo.Login(email, password)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return user.User{}, errors.New("data tidak ditemukan")
		}
		return user.User{}, errors.New("telah terjadi kesalahan terhadap sistem")
	}

	return result, nil
}

func (us *userService) GetUserByEmail(email string) (user.User, error) {
	result, err := us.repo.GetUserByEmail(email)
	if err != nil {
		return user.User{}, err // Adjust the error handling as needed
	}
	return result, nil
}
