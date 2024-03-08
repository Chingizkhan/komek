package mapper

import (
	"komek/internal/domain"
	"strings"
)

func ConvRolesFromStringToDomain(rolesStr string) domain.Roles {
	roles := strings.Split(rolesStr, ",")
	return ConvRolesToDomain(roles)
}

func ConvRolesToDomain(roles []string) domain.Roles {
	resRoles := make(domain.Roles, 0, len(roles))
	for _, r := range roles {
		resRoles = append(resRoles, domain.Role(r))
	}
	return resRoles
}
