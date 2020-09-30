package migrations

import (
	"errors"

	"github.com/flaviowilker/rentcar/app/domain"
	"gorm.io/gorm"
)

var (
	errRolesNotCreated = errors.New("Roles were not created")
)

func migrationV2(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {

		err := func() error {
			role, err := domain.NewRole("Administrator", "ADMIN")
			if err != nil {
				return errRolesNotCreated
			}

			result := tx.Where("code = ?", role.Code).First(&domain.Role{})
			if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return errRolesNotCreated
			}

			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				if err := tx.Create(role).Error; err != nil {
					return errRolesNotCreated
				}
			}

			return nil
		}()
		if err != nil {
			return errRolesNotCreated
		}

		err = func() error {
			role, err := domain.NewRole("Customer", "CUSTOMER")
			if err != nil {
				return errRolesNotCreated
			}

			result := tx.Where("code = ?", role.Code).First(&domain.Role{})
			if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return errRolesNotCreated
			}

			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				if err := tx.Create(role).Error; err != nil {
					return errRolesNotCreated
				}
			}

			return nil
		}()
		if err != nil {
			return errRolesNotCreated
		}

		return nil
	})

	return err
}
