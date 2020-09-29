package domain

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User Errors
var (
	ErrUserNotFound = errors.New("user not found")

	ErrUserInvalidPassword = errors.New("The password is invalid")
)

// User ...
type User struct {
	gorm.Model
	Name          string `gorm:"type:varchar(255)"`
	Login         string `gorm:"type:varchar(255);unique_index"`
	Password      string `gorm:"type:varchar(255)"`
	Token         string `gorm:"type:varchar(255);unique_index"`
	RoleID        uint
	Role          Role
	UserSchedules []UserSchedule
}

// NewUser ...
func NewUser(name string, login string, password string, roleID uint) (*User, error) {
	user := &User{
		Name:     name,
		Login:    login,
		Password: password,
		RoleID:   roleID,
	}

	passwordCrypted, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(passwordCrypted)

	newToken, err := bcrypt.GenerateFromPassword([]byte(user.Login+"-"+user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Token = string(newToken)

	return user, nil
}

// EqualsPassword ...
func (user *User) EqualsPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	return err == nil
}
