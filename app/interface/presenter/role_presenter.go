package presenter

import (
	"github.com/flaviowilker/rentcar/app/application/usecase/input"
	"github.com/flaviowilker/rentcar/app/application/usecase/output"
	"github.com/flaviowilker/rentcar/app/domain"
)

// NewRolePresenter ...
func NewRolePresenter() *RolePresenter {
	return &RolePresenter{}
}

// RolePresenter ...
type RolePresenter struct {
}

// InputRole ...
func (r *RolePresenter) InputRole(input *input.Role) *domain.Role {
	return &domain.Role{
		Name: input.Name,
		Code: input.Code,
	}
}

// OutputRole ...
func (r *RolePresenter) OutputRole(role *domain.Role) *output.Role {
	return &output.Role{
		ID:   role.ID,
		Name: role.Name,
		Code: role.Code,
	}
}

// OutputRoles ...
func (r *RolePresenter) OutputRoles(roles []*domain.Role) []*output.Role {

	outputRoles := make([]*output.Role, len(roles))

	for i, role := range roles {
		outputRoles[i] = r.OutputRole(role)
	}

	return outputRoles
}
