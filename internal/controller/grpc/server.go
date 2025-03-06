package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"komek/internal/controller/grpc/middleware"
	"komek/internal/usecase/user"
	"komek/pb"
	"komek/pkg/logger"
)

type Server struct {
	user *user.UseCase
	l    logger.ILogger
	pb.UnimplementedKomekServer
}

func Register(l logger.ILogger, user *user.UseCase) *grpc.Server {
	server := &Server{
		user: user,
		l:    l,
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.Log(l)),
		//grpc.UnaryInterceptor(middleware.IP),
	)

	pb.RegisterKomekServer(grpcServer, server)
	reflection.Register(grpcServer)

	return grpcServer
}
