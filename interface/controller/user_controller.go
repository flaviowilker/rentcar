package controller

import (
	"github.com/flaviowilker/rentcar/domain"
	"github.com/flaviowilker/rentcar/interface/usecase"
)

// UserController ...
type UserController struct {
	UserUseCase usecase.UserUseCase
}

// NewUserController ...
func NewUserController(u usecase.UserUseCase) UserController {
	return UserController{
		UserUseCase: u,
	}
}

// FindAll ...
func (u *UserController) FindAll() (*[]domain.User, error) {
	users, err := u.UserUseCase.FindAll()

	if err != nil {
		return users, err
	}

	return users, nil
}
