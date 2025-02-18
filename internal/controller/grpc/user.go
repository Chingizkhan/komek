package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"komek/internal/domain/password"
	"komek/internal/domain/phone"
	"komek/internal/dto"
	"komek/internal/mapper"
	"komek/pb"
	"log"
)

func (s *Server) RegisterUser(ctx context.Context, r *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	req := dto.UserRegisterRequest{
		Login:    r.Login,
		Phone:    phone.Phone(r.Phone),
		Password: password.Password(r.Password),
		Roles:    mapper.ConvRolesToDomain(r.Roles),
	}
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	metadata := s.extractMetadata(ctx)
	log.Println("metadata:", metadata)
	user, err := s.user.Register(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.RegisterUserResponse{
		User: mapper.ConvUserPb(user),
	}, nil
}

func (s *Server) LoginUser(ctx context.Context, request *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	//TODO implement me
	panic("implement me")
}
