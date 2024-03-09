package middleware

import (
	"context"
	"google.golang.org/grpc"
	"komek/pkg/logger"
)

func Log(l logger.ILogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			l.Error("error from logger middleware", logger.Err(err))
		}

		return resp, err
	}
}
