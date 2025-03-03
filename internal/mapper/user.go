package mapper

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"komek/db/sqlc"
	"komek/internal/domain/email"
	"komek/internal/domain/phone"
	"komek/internal/domain/user/entity"
	"komek/pb"
	"komek/pkg/money"
)

func ConvUserToDomain(user sqlc.User) entity.User {
	return entity.User{
		ID:                user.ID.Bytes,
		Name:              user.Name.String,
		Phone:             phone.Phone(user.Phone.String),
		Login:             user.Login.String,
		Email:             email.Email(user.Email.String),
		EmailVerified:     user.EmailVerified.Bool,
		PasswordHash:      user.PasswordHash,
		Roles:             ConvRolesFromStringToDomain(user.Roles),
		CreatedAt:         user.CreatedAt.Time,
		UpdatedAt:         user.UpdatedAt.Time,
		PasswordChangedAt: user.PasswordChangedAt.Time,
	}
}

func ConvUserResponse(out entity.GetOut) entity.UserResponse {
	return entity.UserResponse{
		ID:                out.User.ID,
		Name:              out.User.Name,
		Login:             out.User.Login,
		Phone:             out.User.Phone,
		Email:             out.User.Email,
		EmailVerified:     out.User.EmailVerified,
		Roles:             out.User.Roles,
		CreatedAt:         out.User.CreatedAt.Unix(),
		UpdatedAt:         out.User.UpdatedAt.Unix(),
		PasswordChangedAt: out.User.PasswordChangedAt.Unix(),
		Account: entity.AccountResponse{
			ID:       out.Account.ID,
			Balance:  money.ToFloat(out.Account.Balance),
			Currency: out.Account.Currency,
			Country:  out.Account.Country,
		},
	}
}

func ConvUserPb(user entity.User) *pb.User {
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
