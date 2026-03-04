package main

import (
	"context"
	"log"
	"net"
	"os/signal"
	"syscall"
	"time"

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

	serveErrCh := make(chan error, 1)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			serveErrCh <- err
		}
	}()

	signalCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case <-signalCtx.Done():
		log.Println("shutdown signal received")
	case err := <-serveErrCh:
		log.Fatal(err)
	}

	stopped := make(chan struct{})
	go func() {
		grpcServer.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
		log.Println("gRPC server stopped gracefully")
	case <-time.After(10 * time.Second):
		log.Println("gRPC graceful stop timeout, forcing stop")
		grpcServer.Stop()
	}
}
