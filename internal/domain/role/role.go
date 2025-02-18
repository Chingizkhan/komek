package role

import "strings"

type Role string
type Roles []Role

const (
	User    Role = "user"
	Admin   Role = "admin"
	Manager Role = "manager"
)

var roles = Roles{User, Admin, Manager}

func (r Role) IsUser() bool {
	return r == User
}

func (r Role) IsAdmin() bool {
	return r == Admin
}

func (r Role) IsManager() bool {
	return r == Manager
}

func (rs Roles) ToString() string {
	rolesStr := make([]string, 0, len(rs))

	for _, r := range rs {
		rolesStr = append(rolesStr, string(r))
	}

	return strings.Join(rolesStr, ",")
}

func (rs Roles) Allowed() bool {
	var allowed = true
	for _, r := range rs {
		if !roles.Contains(r) {
			allowed = false
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
