package output

// User ...
type User struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Login  string `json:"login"`
	RoleID uint   `json:"roleId"`
}
