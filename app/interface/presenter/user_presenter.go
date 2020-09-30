package presenter

import (
	"github.com/flaviowilker/rentcar/app/application/usecase/input"
	"github.com/flaviowilker/rentcar/app/application/usecase/output"
	"github.com/flaviowilker/rentcar/app/domain"
)

// UserPresenter ...
type UserPresenter struct {
}

// NewUserPresenter ...
func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

// InputUser ...
func (u *UserPresenter) InputUser(input *input.User) *domain.User {
	return &domain.User{
		Name:     input.Name,
		Login:    input.Login,
		Password: input.Password,
		RoleID:   input.RoleID,
	}
}

// OutputUser ...
func (u *UserPresenter) OutputUser(user *domain.User) *output.User {
	return &output.User{
		ID:     user.ID,
		Name:   user.Name,
		Login:  user.Login,
		RoleID: user.RoleID,
	}
}

// OutputUsers ...
func (u *UserPresenter) OutputUsers(users []*domain.User) []*output.User {

	outputUsers := make([]*output.User, len(users))

	for i, user := range users {
		outputUsers[i] = &output.User{
			ID:     user.ID,
			Name:   user.Name,
			Login:  user.Login,
			RoleID: user.RoleID,
		}
	}

	return outputUsers
}
