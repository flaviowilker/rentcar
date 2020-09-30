package migrations

import (
	"log"

	"gorm.io/gorm"
)

var (
	messageErrorMigrations = "Error to execute database migrations"
)

// ExecuteMigrations ...
func ExecuteMigrations(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {

		if err := migrationV1(tx); err != nil {
			return err
		}

		if err := migrationV2(tx); err != nil {
			return err
		}

		if err := migrationV3(tx); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalf("%v: %v", messageErrorMigrations, err)
		panic(err)
	}
}
