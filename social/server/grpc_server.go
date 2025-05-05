package server

import (
	"context"
	"log"
	"net"

	pb "github.com/kzemlyak/microchelik/social/pkg/api/social"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedSocialServiceServer
}

func newServer() *server {
	return &server{}
}

func (s *server) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("Заголовков нет")
	} else {
		const key = "x-header"
		log.Println(key, md.Get(key))
	}

	return &pb.CreateProfileResponse{
		UserId: "asf-asf-asf-asf",
	}, nil
}

func InitGrpcServer() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	implementation := newServer()

	server := grpc.NewServer()
	pb.RegisterSocialServiceServer(server, implementation)

	reflection.Register(server)

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
