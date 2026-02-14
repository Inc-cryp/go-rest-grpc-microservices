package usecase

import (
	"context"

	"github.com/abdillahfazri/grpc-simple/internal/user/domain"
	"github.com/abdillahfazri/grpc-simple/internal/user/repository"
)

type UserUsecase interface {
	GetUser(ctx context.Context, id string) (*domain.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) GetUser(
	ctx context.Context,
	id string,
) (*domain.User, error) {
	return u.repo.FindByID(ctx, id)
}
