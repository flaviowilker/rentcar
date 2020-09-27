package application

import (
	"github.com/flaviowilker/rentcar/domain"
)

// UserUseCase ...
type UserUseCase struct {
	UserRepository domain.UserRepository
}

// NewUserUseCase ...
func NewUserUseCase(u domain.UserRepository) UserUseCase {
	return UserUseCase{
		UserRepository: u,
	}
}

// FindAll ...
func (u *UserUseCase) FindAll() (*[]domain.User, error) {
	users, err := u.UserRepository.FindAll()

	if err != nil {
		return users, err
	}

	return users, nil
}

// Create ...
func (u *UserUseCase) Create(entity domain.User) (*domain.User, error) {

	user, err := u.UserRepository.Create(entity)

	if err != nil {
		return user, err
	}

	return user, nil
}

// Update ...
func (u *UserUseCase) Update(entity domain.User) (*domain.User, error) {

	user, err := u.UserRepository.Update(entity)

	if err != nil {
		return user, err
	}

	return user, nil
}

// Delete ...
func (u *UserUseCase) Delete(id uint) (*domain.User, error) {
	user, err := u.UserRepository.Delete(id)

	if err != nil {
		return user, err
	}

	return user, nil
}

// Login ...
func (u *UserUseCase) Login(login string, password string) (*domain.User, error) {

	user, err := u.UserRepository.FindByLogin(login)

	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	if user.EqualsPassword(password) {
		return user, nil
	}

	return nil, domain.ErrUserInvalidPassword
}
