package output

import uuid "github.com/satori/go.uuid"

// Role ...
type Role struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Code string    `json:"code"`
}
