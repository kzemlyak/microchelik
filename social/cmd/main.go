package main

import (
	"context"
	"log"
	"time"

	controllers "github.com/kzemlyak/microchelik/social/internal/app/controllers"
	friend_repository "github.com/kzemlyak/microchelik/social/internal/app/repositories/friend"
	friend_request_repository "github.com/kzemlyak/microchelik/social/internal/app/repositories/friend_request"
	profile_repository "github.com/kzemlyak/microchelik/social/internal/app/repositories/profile"
	usecases "github.com/kzemlyak/microchelik/social/internal/app/usecases"
	grpc_middleware "github.com/kzemlyak/microchelik/social/internal/middleware/grpc"
	server "github.com/kzemlyak/microchelik/social/internal/server"
	postgres "github.com/kzemlyak/microchelik/social/pkg/postgres"
	transaction_manager "github.com/kzemlyak/microchelik/social/pkg/postgres/transaction_manager"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// =========================
	// adapters
	// =========================

	// repository
	pgConn, err := postgres.NewConnectionPool(ctx, "postgres://jack:secret@pg.example.com:5432/mydb?sslmode=disabled",
		postgres.WithMaxConnIdleTime(time.Minute),
	)
	if err != nil {
		log.Fatal(err)
	}

	txManager := transaction_manager.New(pgConn)

	friendRepo := friend_repository.NewFriendRepository(txManager)
	friendRequestRepo := friend_request_repository.NewFriendRequestRepository(txManager)
	profileRepo := profile_repository.NewProfileRepository(txManager)

	// services
	_, err = grpc.NewClient("localhost:8082",
		grpc.WithUnaryInterceptor(
			grpc_middleware.PropagationUnaryClientInterceptor(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// =========================
	// usecases
	// =========================
	useCases := usecases.NewUseCase(usecases.Deps{
		FriendRepository:        friendRepo,
		FriendRequestRepository: friendRequestRepo,
		ProfileRepository:       profileRepo,
		TxManager:               txManager,
	})

	// =========================
	// delivery
	// =========================
	// controllers
	ctrls := controllers.NewController(controllers.Deps{
		UseCase: useCases,
	})

	// middlewares
	validator, err := protovalidate.New(protovalidate.WithDisableLazy(false))
	if err != nil {
		log.Fatalf("server: failed to initialize validator: %s", err)
	}
	mws := []grpc.UnaryServerInterceptor{
		grpc_middleware.ErrorsUnaryServerInterceptor(),
		grpc_middleware.ValidateUnaryServerInterceptor(validator),
	}

	// infrastructure server
	config := server.Config{
		GRPCPort:               ":8082",
		GRPCGatewayPort:        ":8080",
		ChainUnaryInterceptors: mws,
	}

	srv, err := server.NewServer(ctx, config, server.Controllers{
		SocialServiceServer: ctrls,
	})
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	if err = srv.Run(ctx); err != nil {
		log.Fatalf("run: %v", err)
	}

	server.InitHttpServer()
}
