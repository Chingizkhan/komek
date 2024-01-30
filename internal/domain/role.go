package domain

import "strings"

type Role string
type Roles []Role

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

func (roles Roles) ConvString() string {
	rolesStr := make([]string, 0, len(roles))

	for _, r := range roles {
		rolesStr = append(rolesStr, string(r))
	}

	return strings.Join(rolesStr, ",")
}
