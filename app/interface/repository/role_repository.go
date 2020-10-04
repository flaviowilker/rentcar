package repository

import (
	"github.com/flaviowilker/rentcar/app/domain"
	"gorm.io/gorm"
)

// NewRoleRepository ...
func NewRoleRepository(db *gorm.DB) RoleRepository {
	return RoleRepository{Db: db}
}

// RoleRepository ...
type RoleRepository struct {
	Db *gorm.DB
}

// FindByCode ...
func (r *RoleRepository) FindByCode(code string) (*domain.Role, error) {
	role := new(domain.Role)

	err := r.Db.First(role, "code = ?", code).Error
	if err != nil {
		return nil, err
	}

	return role, nil
}
