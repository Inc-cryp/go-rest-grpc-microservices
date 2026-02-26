package main

import (
	"log"
	"net"

	config "github.com/abdillahfazri/grpc-simple/internal/config"
	"github.com/abdillahfazri/grpc-simple/internal/user/repository"
	unaryLogger "github.com/abdillahfazri/grpc-simple/internal/user/transport/grpc"
	usergrpc "github.com/abdillahfazri/grpc-simple/internal/user/transport/grpc"
	"github.com/abdillahfazri/grpc-simple/internal/user/usecase"
	userpb "github.com/abdillahfazri/grpc-simple/proto/user/v1"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	userGRPCPort := cfg.UserGrpcPort
	if userGRPCPort == "" {
		userGRPCPort = "50051"
	}

	listenAddr := ":" + userGRPCPort

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc.UnaryServerInterceptor(unaryLogger.UnaryLogger)),
	)

	repo := repository.NewUserRepository()
	uc := usecase.NewUserUsecase(repo)
	handler := usergrpc.NewHandler(uc)

	userpb.RegisterUserServiceServer(grpcServer, handler)

	log.Printf("gRPC server running on %s (env=%s)", listenAddr, cfg.AppEnv)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
