package middleware_grpc

import (
	"context"

	grpcutils "lesson_08/orders_management_system/pkg/grpc_utils"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// ValidateUnaryServerInterceptor - convert any arror to rpc error
func ValidateUnaryServerInterceptor(validator *protovalidate.Validator) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		if msg, ok := (req).(proto.Message); ok {
			if err := validator.Validate(msg); err != nil {
				return nil, grpcutils.RPCValidationError(err)
			}
		}

		return handler(ctx, req)
	}
}
