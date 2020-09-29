package usecase

import "github.com/flaviowilker/rentcar/domain"

// UserUseCase ...
type UserUseCase interface {
	FindAll() (*[]domain.User, error)
	Create(domain.User) (*domain.User, error)
	Update(domain.User) (*domain.User, error)
	Delete(uint) (*domain.User, error)
	Login(string, string) (*domain.User, error)
}
