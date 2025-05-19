package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/kzemlyak/microchelik/social/pkg/api/social"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	GRPCPort        string
	GRPCGatewayPort string

	ChainUnaryInterceptors []grpc.UnaryServerInterceptor
	UnaryInterceptors      []grpc.UnaryServerInterceptor
}

type Controllers struct {
	pb.SocialServiceServer
}

type Server struct {
	Controllers

	grpc struct {
		lis    net.Listener
		server *grpc.Server
	}
}

func NewServer(ctx context.Context, cfg Config, ctrls Controllers) (*Server, error) {
	srv := &Server{Controllers: ctrls}

	grpcServerOptions := unaryInterceptorsToGrpcServerOptions(cfg.UnaryInterceptors...)
	grpcServerOptions = append(grpcServerOptions,
		grpc.ChainUnaryInterceptor(cfg.ChainUnaryInterceptors...),
	)

	grpcServer := grpc.NewServer(grpcServerOptions...)
	pb.RegisterSocialServiceServer(grpcServer, srv)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		return nil, fmt.Errorf("server: failed to listen: %v", err)
	}

	srv.grpc.server = grpcServer
	srv.grpc.lis = lis

	return srv, nil
}

func (s *Server) Run(ctx context.Context) error {
	if err := s.grpc.server.Serve(s.grpc.lis); err != nil {
		return fmt.Errorf("server: serve grpc: %v", err)
	}
	return nil
}

func unaryInterceptorsToGrpcServerOptions(interceptors ...grpc.UnaryServerInterceptor) []grpc.ServerOption {
	opts := make([]grpc.ServerOption, 0, len(interceptors))
	for _, interceptor := range interceptors {
		opts = append(opts, grpc.UnaryInterceptor(interceptor))
	}
	return opts
}
