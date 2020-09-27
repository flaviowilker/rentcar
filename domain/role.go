package domain

import "gorm.io/gorm"

// Role ...
type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);unique"`
	Code string `gorm:"type:varchar(255);unique"`
}
