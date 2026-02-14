package main

import (
	"context"
	"log"
	"time"

	userpb "github.com/abdillahfazri/grpc-simple/proto/user/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. Dial ke gRPC server
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 2. Buat STUB (client)
	client := userpb.NewUserServiceClient(conn)

	// 3. Context + timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 4. Call RPC
	resp, err := client.GetUser(ctx, &userpb.GetUserRequest{
		Id: "1",
	})
	if err != nil {
		log.Fatal(err)
	}

	// 5. Print response
	log.Println("Response from gRPC:")
	log.Println("ID   :", resp.Id)
	log.Println("Name :", resp.Name)
	log.Println("Email:", resp.Email)
}
