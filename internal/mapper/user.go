package mapper

import (
	"komek/db/sqlc"
	"komek/internal/domain"
	"strings"
)

func ConvUserToDomain(user sqlc.User) domain.User {
	roles := strings.Split(user.Roles, ",")
	resRoles := make(domain.Roles, 0, len(roles))

	for _, r := range roles {
		resRoles = append(resRoles, domain.Role(r))
	}

	return domain.User{
		ID:            user.ID.Bytes,
		Name:          user.Name.String,
		Phone:         domain.Phone(user.Phone.String),
		Login:         user.Login,
		Email:         domain.Email(user.Email.String),
		EmailVerified: user.EmailVerified.Bool,
		PasswordHash:  user.PasswordHash,
		Roles:         resRoles,
		CreatedAt:     user.CreatedAt.Time,
		UpdatedAt:     user.UpdatedAt.Time,
	}
}
