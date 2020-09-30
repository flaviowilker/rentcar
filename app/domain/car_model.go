package domain

import "gorm.io/gorm"

// CarModel ...
type CarModel struct {
	gorm.Model
	CarModel string `gorm:"type:varchar(255)"`
	Maker    string `gorm:"type:varchar(255)"`
	Color    string `gorm:"type:varchar(255)"`
}
