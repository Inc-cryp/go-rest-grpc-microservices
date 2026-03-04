package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	config "github.com/abdillahfazri/grpc-simple/internal/config"
	orderClient "github.com/abdillahfazri/grpc-simple/internal/order/client"
	orderRest "github.com/abdillahfazri/grpc-simple/internal/order/transport/rest"
	orderUsecase "github.com/abdillahfazri/grpc-simple/internal/order/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	userClient, err := orderClient.NewUserClient(cfg.UserGrpcAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := userClient.Close(); err != nil {
			log.Println("failed to close user gRPC client:", err)
		}
	}()

	uc := orderUsecase.NewOrderUsecase(userClient)
	handler := orderRest.NewHandler(uc)

	r.POST("/orders/:user_id", handler.CreateOrder)

	orderServer := &http.Server{
		Addr:    ":" + cfg.OrderHTTPPort,
		Handler: r,
	}
	serverErrCh := make(chan error, 1)

	log.Println("Order API running on :" + cfg.OrderHTTPPort)
	go func() {
		if err := orderServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrCh <- err
		}
	}()

	signalCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case <-signalCtx.Done():
		log.Println("shutdown signal received")
	case err := <-serverErrCh:
		log.Println("server error:", err)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := orderServer.Shutdown(shutdownCtx); err != nil {
		log.Println("order server shutdown error:", err)
	}

}
