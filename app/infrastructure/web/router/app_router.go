package router

import (
	"github.com/flaviowilker/rentcar/app/application/usecase"
	"github.com/flaviowilker/rentcar/app/interface/controller"
	"github.com/flaviowilker/rentcar/app/interface/presenter"
	"github.com/flaviowilker/rentcar/app/interface/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"gorm.io/gorm"
)

// CreateAppRouter ...
func CreateAppRouter(db *gorm.DB, app *fiber.App) {
	appRouter := &appRouter{
		db:  db,
		app: app,
	}

	appRouter.createServerEndpoints()
	appRouter.createUserEndpoints()
}

type appRouter struct {
	db  *gorm.DB
	app *fiber.App
}

func (a *appRouter) createServerEndpoints() {
	a.app.Get("/dashboard", monitor.New())
}

func (a *appRouter) createUserEndpoints() {
	roleRepository := repository.NewRoleRepository(a.db)
	userRepository := repository.NewUserRepository(a.db)
	userUseCase := usecase.NewUserUseCase(&userRepository, presenter.NewUserPresenter(), &roleRepository, presenter.NewRolePresenter())
	userController := controller.NewUserController(&userUseCase)
	userRouter := newUserRouter(a.app, &userController)
	userRouter.createRoutes()
}
