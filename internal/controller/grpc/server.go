package grpc

import (
	"google.golang.org/grpc"
	"komek/internal/controller/grpc/middleware"
	"komek/pb"
	"komek/pkg/logger"
)

type Server struct {
	l logger.ILogger
	pb.UnimplementedKomekServer
}

func Register(l logger.ILogger) *grpc.Server {
	server := &Server{l: l}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.Log(l)),
		//grpc.UnaryInterceptor(middleware.IP),
	)

	pb.RegisterKomekServer(grpcServer, server)

	return grpcServer
}
