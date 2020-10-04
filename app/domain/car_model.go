package domain

// CarModel ...
type CarModel struct {
	Base
	CarModel string `gorm:"type:varchar(255)"`
	Maker    string `gorm:"type:varchar(255)"`
	Color    string `gorm:"type:varchar(255)"`
}
