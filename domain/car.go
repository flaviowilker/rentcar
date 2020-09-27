package domain

import "gorm.io/gorm"

// Car ...
type Car struct {
	gorm.Model
	LicensePlate string `gorm:"type:varchar(255);unique"`
	CarModelID   uint
	CarModel     CarModel
}
