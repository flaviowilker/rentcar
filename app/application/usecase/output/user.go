package output

import uuid "github.com/satori/go.uuid"

// User ...
type User struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Login  string    `json:"login"`
	RoleID uuid.UUID `json:"roleId"`
}
