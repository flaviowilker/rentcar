package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Base ...
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate ...
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.NewV4()
	return nil
}
