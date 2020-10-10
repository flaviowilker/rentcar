package router

import (
	"github.com/flaviowilker/rentcar/app/infrastructure/controller"
	"github.com/gofiber/fiber/v2"
)

func newUserRouter(a *fiber.App, u controller.UserController) userRouter {
	return userRouter{
		app:            a,
		userController: u,
	}
}

type userRouter struct {
	app            *fiber.App
	userController controller.UserController
}

func (u *userRouter) createRoutes() {
	u.app.Get("/v1/users", u.userController.FindAll)
	u.app.Post("/v1/users", u.userController.Create)
}
