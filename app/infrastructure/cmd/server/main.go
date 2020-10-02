package main

import (
	"errors"
	"log"
	"os"

	"github.com/flaviowilker/rentcar/app/application/usecase"
	"github.com/flaviowilker/rentcar/app/infrastructure/database"
	"github.com/flaviowilker/rentcar/app/infrastructure/router"
	"github.com/flaviowilker/rentcar/app/interface/controller"
	"github.com/flaviowilker/rentcar/app/interface/presenter"
	"github.com/flaviowilker/rentcar/app/interface/repository"

	fiber "github.com/gofiber/fiber/v2"
)

var (
	errPortIsNull = errors.New("Error web port is null")
)

func main() {
	port := os.Getenv("WEB_PORT")
	if port == "" {
		log.Fatal(errPortIsNull)
	}

	db := database.NewConfig()
	dbSQL, err := db.DB()
	if err != nil {
		defer dbSQL.Close()
	}

	app := fiber.New()

	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(&userRepository, presenter.NewUserPresenter())
	userController := controller.NewUserController(&userUseCase)
	userRouter := router.NewUserUserRouter(app, &userController)
	userRouter.CreateRoutes()

	log.Fatal(app.Listen(":" + port))
}
