package controller

import (
	"github.com/flaviowilker/rentcar/app/application/usecase/output"
	"github.com/flaviowilker/rentcar/app/interface/usecase"
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
func (u *UserController) FindAll() ([]*output.User, error) {
	users, err := u.UserUseCase.FindAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}
