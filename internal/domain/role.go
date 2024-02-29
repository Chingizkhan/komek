package domain

import "strings"

type Role string
type Roles []Role

const (
	RoleUser    Role = "user"
	RoleAdmin   Role = "admin"
	RoleManager Role = "manager"
)

var roles = Roles{RoleUser, RoleAdmin, RoleManager}

func (r Role) IsUser() bool {
	return r == RoleUser
}

func (r Role) IsAdmin() bool {
	return r == RoleAdmin
}

func (r Role) IsManager() bool {
	return r == RoleManager
}

func (rs Roles) ConvString() string {
	rolesStr := make([]string, 0, len(rs))

	for _, r := range rs {
		rolesStr = append(rolesStr, string(r))
	}

	return strings.Join(rolesStr, ",")
}

func (rs Roles) Allowed() bool {
	var allowed bool
	for _, r := range rs {
		if roles.Contains(r) {
			allowed = true
		}
	}
	return allowed
}

func (rs Roles) Contains(role Role) bool {
	var exists bool
	for _, r := range roles {
		if r == role {
			exists = true
		}
	}
	return exists
}
