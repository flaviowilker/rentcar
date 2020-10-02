package controller

import (
	"github.com/flaviowilker/rentcar/app/interface/usecase"
	"github.com/gofiber/fiber/v2"
)

// NewUserController ...
func NewUserController(u usecase.UserUseCase) UserController {
	return UserController{
		UserUseCase: u,
	}
}

// UserController ...
type UserController struct {
	UserUseCase usecase.UserUseCase
}

// FindAll ...
func (u *UserController) FindAll(c *fiber.Ctx) error {
	users, err := u.UserUseCase.FindAll()

	if err != nil {
		return c.SendStatus(404)
	}

	return c.JSON(users)
}
