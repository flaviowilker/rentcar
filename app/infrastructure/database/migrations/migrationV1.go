package migrations

import (
	"errors"

	"github.com/flaviowilker/rentcar/app/domain"
	"gorm.io/gorm"
)

var (
	errTablesNotCreated = errors.New("Tables were not created")
)

func migrationV1(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {

		err := tx.AutoMigrate(&domain.Role{},
			&domain.CarModel{},
			&domain.Car{},
			&domain.CarSchedule{},
			&domain.User{},
			&domain.UserSchedule{})

		if err != nil {
			return errTablesNotCreated
		}

		return nil
	})

	return err
}
