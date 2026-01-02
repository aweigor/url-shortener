package auth_test

import (
	"testing"
	"url-shortener/internal/auth"
	"url-shortener/internal/user"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
func (repo *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	const originalEmail = "test@test.test"
	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register(originalEmail, "password", "name")
	if err != nil {
		t.Fatal(err)
	}
	if email != originalEmail {
		t.Fatalf("Email %s not match %s", email, originalEmail)
	}
}
