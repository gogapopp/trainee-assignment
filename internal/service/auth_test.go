package service

import (
	"context"
	"testing"

	"github.com/gogapopp/trainee-assignment/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthRepo struct {
	mock.Mock
}

func (m *MockAuthRepo) SignUp(ctx context.Context, user models.SignUpRequest) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockAuthRepo) SignIn(ctx context.Context, user models.SignInRequest) (string, error) {
	args := m.Called(ctx, user)
	return args.String(0), args.Error(1)
}

func TestAuthService_SignUp(t *testing.T) {
	authRepo := new(MockAuthRepo)
	s := NewAuthService("jwtSecret", "passSecret", nil, authRepo)

	tests := []struct {
		name string
		user models.SignUpRequest
		mock func()
	}{
		{
			name: "#1 ok",
			user: models.SignUpRequest{Username: "admin", PasswordHash: "pass", Role: "admin"},
			mock: func() {
				authRepo.On("SignUp", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "#2 ok",
			user: models.SignUpRequest{Username: "user", PasswordHash: "pass", Role: "user"},
			mock: func() {
				authRepo.On("SignUp", mock.Anything, mock.Anything).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := s.SignUp(context.Background(), tt.user)
			assert.NoError(t, err)
		})
	}
}

func TestAuthService_SignIn(t *testing.T) {
	authRepo := new(MockAuthRepo)
	s := NewAuthService("jwtSecret", "passSecret", nil, authRepo)

	tests := []struct {
		name string
		user models.SignInRequest
		mock func()
	}{
		{
			name: "#1 ok",
			user: models.SignInRequest{Username: "user", PasswordHash: "hash"},
			mock: func() {
				authRepo.On("SignIn", mock.Anything, mock.Anything).Return("token", nil)
			},
		},
		{
			name: "#2 ok",
			user: models.SignInRequest{Username: "admin", PasswordHash: "hash"},
			mock: func() {
				authRepo.On("SignIn", mock.Anything, mock.Anything).Return("token", nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := s.SignIn(context.Background(), tt.user)
			assert.NoError(t, err)
			assert.NotEmpty(t, got)
		})
	}
}
