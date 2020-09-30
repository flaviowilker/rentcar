package main

import (
	"fmt"

	"github.com/flaviowilker/rentcar/app/application/usecase"
	"github.com/flaviowilker/rentcar/app/infrastructure/database"
	"github.com/flaviowilker/rentcar/app/interface/controller"
	"github.com/flaviowilker/rentcar/app/interface/presenter"
	"github.com/flaviowilker/rentcar/app/interface/repository"
)

func main() {
	db := database.NewConfig()
	dbSQL, err := db.DB()
	if err != nil {
		defer dbSQL.Close()
	}

	// test
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(&userRepository, presenter.NewUserPresenter())
	userController := controller.NewUserController(&userUseCase)

	users, _ := userController.FindAll()
	for _, user := range users {
		fmt.Println(*user)
	}
}
