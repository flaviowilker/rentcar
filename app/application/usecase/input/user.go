package input

// User ...
type User struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	RoleID   uint   `json:"roleId"`
}
