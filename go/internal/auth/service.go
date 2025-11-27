package auth

import (
	"errors"
	"url-shortener/internal/user"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	exitedUser, _ := service.UserRepository.FindByEmail(email)
	if exitedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	user := &user.User{
		Email: email,
		Password: "",
		Name: name,
	}
	_,err := service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}