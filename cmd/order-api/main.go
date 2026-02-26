package main

import (
	"log"

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

	log.Println("Order API running on :" + cfg.OrderHTTPPort)
	r.Run(":" + cfg.OrderHTTPPort)

}
