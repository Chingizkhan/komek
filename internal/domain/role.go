package domain

type Role string

const (
	RoleUser    Role = "user"
	RoleAdmin   Role = "admin"
	RoleManager Role = "manager"
)

func (r Role) IsUser() bool {
	return r == RoleUser
}

func (r Role) IsAdmin() bool {
	return r == RoleAdmin
}

func (r Role) IsManager() bool {
	return r == RoleManager
}
