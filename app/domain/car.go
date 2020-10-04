package domain

// Car ...
type Car struct {
	Base
	LicensePlate string `gorm:"type:varchar(255);unique"`
	CarModelID   uint
	CarModel     CarModel
}
