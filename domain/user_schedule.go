package domain

import (
	"gorm.io/gorm"
)

// UserSchedule ...
type UserSchedule struct {
	gorm.Model
	UserID        uint
	CarScheduleID uint
	CarSchedule   CarSchedule
}
