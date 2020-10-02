package controller

import fiber "github.com/gofiber/fiber/v2"

// UserController ...
type UserController interface {
	FindAll(*fiber.Ctx) error
}
