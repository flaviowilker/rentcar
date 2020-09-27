package repository

import (
	"github.com/flaviowilker/rentcar/domain"
	"gorm.io/gorm"
)

// UserRepository ...
type UserRepository struct {
	Db *gorm.DB
}

// NewUserRepository ...
func NewUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{Db: DB}
}

// Login ...
func (u *UserRepository) Login(login string) (*domain.User, error) {

	var user domain.User
	u.Db.First(&user, "login = ?", login)

	return &user, nil
}

// FindAll ...
func (u *UserRepository) FindAll() (*[]domain.User, error) {
	var users []domain.User
	u.Db.Find(&users)

	return users, nil
}

// Create ...
func (u *UserRepository) Create(user *domain.User) (*domain.User, error) {
	u.Db.Save(&user)

	return user, nil
}

// Update ...
func (u *UserRepository) Update(user *domain.User) (*domain.User, error) {
	u.Db.Save(&user)

	return user, nil
}

// Delete ...
func (u *UserRepository) Delete(id uint) (*domain.User, error) {
	var user domain.User
	u.Db.First(&user, id)
	u.Db.Delete(&user)

	return &user, nil
}
