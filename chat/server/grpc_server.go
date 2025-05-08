package server

import (
	"context"
	"log"
	"net"

	pb "github.com/kzemlyak/microchelik/chat/pkg/api/chat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func newServer() *server {
	return &server{}
}

func (s *server) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("Заголовков нет")
	} else {
		const key = "x-header"
		log.Println(key, md.Get(key))
	}

	info := req.GetText()

	log.Printf("SendMessage: received: %s", info)

	return &pb.SendMessageResponse{
		Id: "asf-asf-asf-asf",
	}, nil
}

func InitGrpcServer() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	implementation := newServer()

	server := grpc.NewServer()
	pb.RegisterChatServiceServer(server, implementation)

	reflection.Register(server)

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
