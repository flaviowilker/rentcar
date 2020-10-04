package domain

import (
	"time"
)

// CarSchedule ...
type CarSchedule struct {
	Base
	RentDate time.Time `gorm:"type:date"`
	CarID    uint
	Car      Car
}
