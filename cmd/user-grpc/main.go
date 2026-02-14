package main

import (
	"log"
	"net"

	"github.com/abdillahfazri/grpc-simple/internal/user/repository"
	unaryLogger "github.com/abdillahfazri/grpc-simple/internal/user/transport/grpc"
	usergrpc "github.com/abdillahfazri/grpc-simple/internal/user/transport/grpc"
	"github.com/abdillahfazri/grpc-simple/internal/user/usecase"
	userpb "github.com/abdillahfazri/grpc-simple/proto/user/v1"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
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

	log.Println("gRPC server running on :50051")
	grpcServer.Serve(lis)
}
