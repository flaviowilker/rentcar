package presenter

import (
	"github.com/flaviowilker/rentcar/app/application/usecase/input"
	"github.com/flaviowilker/rentcar/app/application/usecase/output"
	"github.com/flaviowilker/rentcar/app/domain"
)

// UserPresenter ...
type UserPresenter interface {
	InputUser(*input.User) *domain.User
	OutputUser(*domain.User) *output.User
	OutputUsers([]*domain.User) []*output.User
}
