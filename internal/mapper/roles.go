package mapper

import (
	"komek/internal/domain/role"
	"strings"
)

func ConvRolesFromStringToDomain(rolesStr string) role.Roles {
	roles := strings.Split(rolesStr, ",")
	return ConvRolesToDomain(roles)
}

func ConvRolesToDomain(roles []string) role.Roles {
	resRoles := make(role.Roles, 0, len(roles))
	for _, r := range roles {
		resRoles = append(resRoles, role.Role(r))
	}
	return resRoles
}
