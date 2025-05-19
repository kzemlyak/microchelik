package middleware_grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// AuthorizationUnaryServerInterceptor - convert any arror to rpc error
func AuthorizationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, err
		}

		_ = md.Get("mt token") // GET TOKEN FROM metadata
		// проверка токе
		return handler(ctx, req)
	}
}
