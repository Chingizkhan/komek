package mapper

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"komek/db/sqlc"
	"komek/internal/domain"
	"komek/internal/dto"
	"komek/pb"
)

func ConvUserToDomain(user sqlc.User) domain.User {
	return domain.User{
		ID:                user.ID.Bytes,
		Name:              user.Name.String,
		Phone:             domain.Phone(user.Phone.String),
		Login:             user.Login,
		Email:             domain.Email(user.Email.String),
		EmailVerified:     user.EmailVerified.Bool,
		PasswordHash:      user.PasswordHash,
		Roles:             ConvRolesFromStringToDomain(user.Roles),
		CreatedAt:         user.CreatedAt.Time,
		UpdatedAt:         user.UpdatedAt.Time,
		PasswordChangedAt: user.PasswordChangedAt.Time,
	}
}

func ConvUserResponse(user domain.User) dto.UserResponse {
	return dto.UserResponse{
		ID:                user.ID,
		Name:              user.Name,
		Login:             user.Login,
		Phone:             user.Phone,
		Email:             user.Email,
		EmailVerified:     user.EmailVerified,
		Roles:             user.Roles,
		CreatedAt:         user.CreatedAt.Unix(),
		UpdatedAt:         user.UpdatedAt.Unix(),
		PasswordChangedAt: user.PasswordChangedAt.Unix(),
	}
}

func ConvUserPb(user domain.User) *pb.User {
	return &pb.User{
		Id:                user.ID.String(),
		Name:              user.Name,
		Login:             user.Login,
		Phone:             string(user.Phone),
		Email:             string(user.Email),
		EmailVerified:     user.EmailVerified,
		CreatedAt:         timestamppb.New(user.CreatedAt),
		UpdatedAt:         timestamppb.New(user.UpdatedAt),
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
	}
}
