package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/abdillahfazri/grpc-simple/internal/user/domain"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

// FindByID mocks the FindByID method of UserRepository
func (m *MockUserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	args := m.Called(ctx, id)
	var user *domain.User
	if args.Get(0) != nil {
		user = args.Get(0).(*domain.User)
	}
	return user, args.Error(1)
}

func TestGetUser(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name         string
		userID       string
		mockRepo     func() *MockUserRepository
		expectedUser *domain.User
		expectedErr  error
	}{
		{
			name:   "Success: User found",
			userID: "user-123",
			mockRepo: func() *MockUserRepository {
				mockRepo := new(MockUserRepository)
				mockRepo.On("FindByID", ctx, "user-123").Return(&domain.User{
					ID:    "user-123",
					Name:  "John Doe",
					Email: "john.doe@example.com",
				}, nil)
				return mockRepo
			},
			expectedUser: &domain.User{
				ID:    "user-123",
				Name:  "John Doe",
				Email: "john.doe@example.com",
			},
			expectedErr: nil,
		},
		{
			name:   "Failure: User not found",
			userID: "user-456",
			mockRepo: func() *MockUserRepository {
				mockRepo := new(MockUserRepository)
				mockRepo.On("FindByID", ctx, "user-456").Return(nil, domain.ErrUserNotFound)
				return mockRepo
			},
			expectedUser: nil,
			expectedErr:  domain.ErrUserNotFound,
		},
		{
			name:   "Failure: Repository error",
			userID: "user-789",
			mockRepo: func() *MockUserRepository {
				mockRepo := new(MockUserRepository)
				mockRepo.On("FindByID", ctx, "user-789").Return(nil, errors.New("database connection error"))
				return mockRepo
			},
			expectedUser: nil,
			expectedErr:  errors.New("database connection error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := tt.mockRepo()
			usecase := NewUserUsecase(mockRepo)

			user, err := usecase.GetUser(ctx, tt.userID)

			assert.Equal(t, tt.expectedUser, user)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
