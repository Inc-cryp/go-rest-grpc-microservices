package grpc

import (
	"context"

	"github.com/abdillahfazri/grpc-simple/internal/user/usecase"
	userpb "github.com/abdillahfazri/grpc-simple/proto/user/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	userpb.UnimplementedUserServiceServer
	uc usecase.UserUsecase
}

func NewHandler(uc usecase.UserUsecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) GetUser(
	ctx context.Context,
	req *userpb.GetUserRequest,
) (*userpb.GetUserResponse, error) {

	user, err := h.uc.GetUser(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &userpb.GetUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
