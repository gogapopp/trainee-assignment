package service

import (
	"context"

	"github.com/gogapopp/trainee-assignment/internal/models"
	"go.uber.org/zap"
)

type (
	Repository interface {
		SignUp(ctx context.Context, user models.SignUpRequest) error
		SignIn(ctx context.Context, user models.SignInRequest) (int, string, error)
	}

	service struct {
		logger     *zap.SugaredLogger
		repo       Repository
		jwtSecret  string
		passSecret string
	}
)

func New(jwtSecret, passSecret string, logger *zap.SugaredLogger, repository Repository) *service {
	return &service{
		logger:     logger,
		repo:       repository,
		jwtSecret:  jwtSecret,
		passSecret: passSecret,
	}
}
