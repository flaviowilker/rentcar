package repository

import (
	"github.com/flaviowilker/rentcar/app/domain"
	"gorm.io/gorm"
)

// NewUserRepository ...
func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{Db: db}
}

// UserRepository ...
type UserRepository struct {
	Db *gorm.DB
}

// FindByLogin ...
func (u *UserRepository) FindByLogin(login string) (*domain.User, error) {
	user := new(domain.User)

	err := u.Db.First(user, "login = ?", login).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindAll ...
func (u *UserRepository) FindAll() ([]*domain.User, error) {
	users := make([]*domain.User, 10)

	err := u.Db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Create ...
func (u *UserRepository) Create(user *domain.User) (*domain.User, error) {
	err := u.Db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update ...
func (u *UserRepository) Update(user *domain.User) (*domain.User, error) {
	err := u.Db.Save(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Delete ...
func (u *UserRepository) Delete(id uint) (*domain.User, error) {
	user := new(domain.User)

	err := u.Db.First(user, id).Error
	if err != nil {
		return nil, err
	}

	err = u.Db.Delete(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
