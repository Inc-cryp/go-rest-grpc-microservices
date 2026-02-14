package usecase

import (
	"context"
	"errors"

	"github.com/abdillahfazri/grpc-simple/internal/order/client"
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
		return "", errors.New("user not found")
	}

	// contoh logic
	return "Order created for user: " + user.Name, nil
}
