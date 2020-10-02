package controller

import (
	"github.com/flaviowilker/rentcar/app/interface/usecase"
	fiber "github.com/gofiber/fiber/v2"
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
func (u *UserController) FindAll(c *fiber.Ctx) error {
	users, err := u.UserUseCase.FindAll()

	if err != nil {
		return c.SendStatus(404)
	}

	return c.JSON(users)
}
