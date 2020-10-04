package usecase

import (
	"github.com/flaviowilker/rentcar/app/application/usecase/input"
	"github.com/flaviowilker/rentcar/app/application/usecase/output"
)

// UserUseCase ...
type UserUseCase interface {
	FindAll() ([]*output.User, error)
	CreateCustomer(*input.User) (*output.User, error)
	Update(*input.User) (*output.User, error)
	Delete(uint) (*output.User, error)
	Login(string, string) (*output.User, error)
}
