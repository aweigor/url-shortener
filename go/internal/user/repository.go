package user

import (
	"url-shortener/pkg/db"
)


type UserRepositoryDeps struct {
	Database *db.Db
}

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	res := repo.Database.DB.First(&user, "email = ?", email)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
