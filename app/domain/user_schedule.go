package domain

// UserSchedule ...
type UserSchedule struct {
	Base
	UserID        uint
	CarScheduleID uint
	CarSchedule   CarSchedule
}
