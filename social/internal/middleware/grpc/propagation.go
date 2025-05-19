package middleware_grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// PropagationUnaryClientInterceptor - интерсептор который прбрасывает выбранные заголовки в следующий сервис
func PropagationUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			var allowedHeaders = []string{
				"authorization",
				// ...
			}
			headers := make(metadata.MD, len(allowedHeaders))
			for _, header := range allowedHeaders {
				if v := md.Get(header); len(v) > 0 {
					headers[header] = v
				}
			}
			ctx = metadata.NewOutgoingContext(ctx, headers)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
