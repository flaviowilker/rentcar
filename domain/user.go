package domain

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User Errors
var (
	ErrUserNotFound = errors.New("user not found")

	ErrUserInvalidPassword = errors.New("The password is invalid")
)

// UserRepository ...
type UserRepository interface {
	FindByLogin(string) (*User, error)
	FindAll() (*[]User, error)
	Create(*User) (*User, error)
	Update(*User) (*User, error)
	Delete(uint) (*User, error)
}

// User ...
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255)"`
	Login    string `gorm:"type:varchar(255);unique_index"`
	Password string `gorm:"type:varchar(255)"`
	Token    string `gorm:"type:varchar(255);unique_index"`
}

// NewUser ...
func NewUser(name string, login string, password string) (*User, error) {
	user := &User{
		Name:     name,
		Login:    login,
		Password: password,
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
