package usecase

import (
	"context"
	"errors"

	"github.com/abdillahfazri/grpc-simple/internal/order/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrUserServiceTimeout    = errors.New("user service timeout")
	ErrUserServiceUnavailable = errors.New("user service unavailable")
	ErrUserServiceFailure    = errors.New("failed to fetch user")
)

type OrderUsecase struct {
	userClient *client.UserClient
}

func NewOrderUsecase(userClient *client.UserClient) *OrderUsecase {
	return &OrderUsecase{userClient: userClient}
}

func (uc *OrderUsecase) CreateOrder(ctx context.Context, userID string) (string, error) {
	user, err := uc.userClient.GetUser(ctx, userID)
	if err != nil {
		if grpcStatus, ok := status.FromError(err); ok {
			switch grpcStatus.Code() {
			case codes.NotFound:
				return "", ErrUserNotFound
			case codes.DeadlineExceeded:
				return "", ErrUserServiceTimeout
			case codes.Unavailable:
				return "", ErrUserServiceUnavailable
			default:
				return "", ErrUserServiceFailure
			}
		}
		return "", ErrUserServiceFailure
	}

	// contoh logic
	return "Order created for user: " + user.Name, nil
}
