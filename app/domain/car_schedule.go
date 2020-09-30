package domain

import (
	"time"

	"gorm.io/gorm"
)

// CarSchedule ...
type CarSchedule struct {
	gorm.Model
	RentDate time.Time `gorm:"type:date"`
	CarID    uint
	Car      Car
}
