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

func (u *UserRouter) createRoutes() {
	u.Router.Get("/v1/users", u.UserController.FindAll)
}

// CreateUserRouter ...
func CreateUserRouter(r *fiber.App, u controller.UserController) {
	userRouter := &UserRouter{
		Router:         r,
		UserController: u,
	}

	userRouter.createRoutes()
}
