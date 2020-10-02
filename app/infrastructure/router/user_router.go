package router

import (
	"github.com/flaviowilker/rentcar/app/infrastructure/controller"
	fiber "github.com/gofiber/fiber/v2"
)

// UserRouter ...
type UserRouter struct {
	Router         *fiber.App
	UserController controller.UserController
}

// NewUserUserRouter ...
func NewUserUserRouter(r *fiber.App, u controller.UserController) UserRouter {
	return UserRouter{
		Router:         r,
		UserController: u,
	}
}

// CreateRoutes ...
func (u *UserRouter) CreateRoutes() {
	u.Router.Get("/v1/users", u.UserController.FindAll)
}
