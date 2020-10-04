package repository

import "github.com/flaviowilker/rentcar/app/domain"

// RoleRepository ...
type RoleRepository interface {
	FindByCode(string) (*domain.Role, error)
}
