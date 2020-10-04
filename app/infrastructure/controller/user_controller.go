package controller

import "github.com/gofiber/fiber/v2"

// UserController ...
type UserController interface {
	FindAll(*fiber.Ctx) error
	Create(c *fiber.Ctx) error
}
