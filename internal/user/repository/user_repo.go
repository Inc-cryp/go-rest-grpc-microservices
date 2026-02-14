package repository

import (
	"context"

	"github.com/abdillahfazri/grpc-simple/internal/user/domain"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*domain.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindByID(
	ctx context.Context,
	id string,
) (*domain.User, error) {

	if id != "1" {
		return nil, domain.ErrUserNotFound
	}

	return &domain.User{
		ID:    "1",
		Name:  "Fazri",
		Email: "fazri@gmail.com",
	}, nil
}
