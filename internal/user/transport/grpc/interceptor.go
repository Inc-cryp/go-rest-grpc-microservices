package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func UnaryLogger(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {

	log.Println("gRPC call:", info.FullMethod)
	return handler(ctx, req)
}
