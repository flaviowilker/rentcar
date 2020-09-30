package migrations

import (
	"errors"

	"github.com/flaviowilker/rentcar/domain"
	"gorm.io/gorm"
)

var (
	errAdminNotCreated = errors.New("Admin user were not created")
)

func migrationV3(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {

		roleAdmin := domain.Role{}
		resultRole := tx.Where("code = ?", "ADMIN").First(&roleAdmin)
		if resultRole.Error != nil && !errors.Is(resultRole.Error, gorm.ErrRecordNotFound) {
			return errAdminNotCreated
		}

		user, err := domain.NewUser("Administrator", "admin", "admin", roleAdmin.ID)
		if err != nil {
			return errAdminNotCreated
		}

		resultUser := tx.Where("login = ?", user.Login).First(&domain.User{})

		if errors.Is(resultUser.Error, gorm.ErrRecordNotFound) {
			if err := tx.Create(user).Error; err != nil {
				return errAdminNotCreated
			}
		}

		return nil
	})

	return err
}
