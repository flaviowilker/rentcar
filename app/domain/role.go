package domain

import "gorm.io/gorm"

// NewRole ...
func NewRole(name string, code string) (*Role, error) {
	role := &Role{
		Name: name,
		Code: code,
	}

	return role, nil
}

// Role ...
type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);unique"`
	Code string `gorm:"type:varchar(255);unique"`
}
