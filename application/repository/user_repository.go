package repository

import "github.com/flaviowilker/rentcar/domain"

// UserRepository ...
type UserRepository interface {
	FindByLogin(string) (*domain.User, error)
	FindAll() (*[]domain.User, error)
	Create(domain.User) (*domain.User, error)
	Update(domain.User) (*domain.User, error)
	Delete(uint) (*domain.User, error)
}
