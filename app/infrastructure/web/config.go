package web

import (
	"errors"
	"log"
	"os"

	"github.com/flaviowilker/rentcar/app/application/usecase"
	"github.com/flaviowilker/rentcar/app/infrastructure/web/router"
	"github.com/flaviowilker/rentcar/app/interface/controller"
	"github.com/flaviowilker/rentcar/app/interface/presenter"
	"github.com/flaviowilker/rentcar/app/interface/repository"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	errPortIsNull = errors.New("Error web port is null")
)

// CreateConfig ...
func CreateConfig(db *gorm.DB) {

	configWeb := &config{os.Getenv("WEB_PORT")}
	if err := configWeb.configValidate(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	app := fiber.New()

	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(&userRepository, presenter.NewUserPresenter())
	userController := controller.NewUserController(&userUseCase)
	router.CreateUserRouter(app, &userController)

	if err := app.Listen(":" + configWeb.port); err != nil {
		log.Fatal(err)
		panic(err)
	}
}

type config struct {
	port string
}

func (c *config) configValidate() error {
	if c.port == "" {
		return errPortIsNull
	}

	return nil
}
