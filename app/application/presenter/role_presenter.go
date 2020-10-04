package presenter

import (
	"github.com/flaviowilker/rentcar/app/application/usecase/input"
	"github.com/flaviowilker/rentcar/app/application/usecase/output"
	"github.com/flaviowilker/rentcar/app/domain"
)

// RolePresenter ...
type RolePresenter interface {
	InputRole(*input.Role) *domain.Role
	OutputRole(*domain.Role) *output.Role
	OutputRoles([]*domain.Role) []*output.Role
}
