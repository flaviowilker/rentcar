package usecase

import (
	"errors"

	"github.com/flaviowilker/rentcar/app/application/presenter"
	"github.com/flaviowilker/rentcar/app/application/repository"
	"github.com/flaviowilker/rentcar/app/application/usecase/input"
	"github.com/flaviowilker/rentcar/app/application/usecase/output"
)

var (
	errUserNotFound        = errors.New("User not found")
	errUserInvalidPassword = errors.New("The password is invalid")
)

// NewUserUseCase ...
func NewUserUseCase(r repository.UserRepository, p presenter.UserPresenter) UserUseCase {
	return UserUseCase{
		UserRepository: r,
		UserPresenter:  p,
	}
}

// UserUseCase ...
type UserUseCase struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

// FindAll ...
func (u *UserUseCase) FindAll() ([]*output.User, error) {
	users, err := u.UserRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return u.UserPresenter.OutputUsers(users), nil
}

// Create ...
func (u *UserUseCase) Create(input *input.User) (*output.User, error) {
	user, err := u.UserRepository.Create(u.UserPresenter.InputUser(input))

	if err != nil {
		return nil, err
	}

	return u.UserPresenter.OutputUser(user), nil
}

// Update ...
func (u *UserUseCase) Update(input *input.User) (*output.User, error) {

	user, err := u.UserRepository.Update(u.UserPresenter.InputUser(input))

	if err != nil {
		return nil, err
	}

	return u.UserPresenter.OutputUser(user), nil
}

// Delete ...
func (u *UserUseCase) Delete(id uint) (*output.User, error) {
	user, err := u.UserRepository.Delete(id)

	if err != nil {
		return nil, err
	}

	return u.UserPresenter.OutputUser(user), nil
}

// Login ...
func (u *UserUseCase) Login(login string, password string) (*output.User, error) {

	user, err := u.UserRepository.FindByLogin(login)

	if err != nil {
		return nil, errUserNotFound
	}

	if user.EqualsPassword(password) {
		return u.UserPresenter.OutputUser(user), nil
	}

	return nil, errUserInvalidPassword
}
