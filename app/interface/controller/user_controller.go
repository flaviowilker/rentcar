package controller

import (
	"github.com/flaviowilker/rentcar/app/application/usecase/input"
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
		return c.Status(404).SendString(err.Error())
	}

	return c.JSON(users)
}

// Create ...
func (u *UserController) Create(c *fiber.Ctx) error {

	userInput := new(input.User)
	if err := c.BodyParser(userInput); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	userOutput, err := u.UserUseCase.CreateCustomer(userInput)

	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.Status(201).JSON(userOutput)
}
