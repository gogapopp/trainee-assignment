package service

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/gogapopp/trainee-assignment/internal/models"
	"go.uber.org/zap"
)

type (
	authRepo interface {
		SignUp(ctx context.Context, user models.SignUpRequest) error
		SignIn(ctx context.Context, user models.SignInRequest) (int, string, error)
	}

	bannerRepo interface {
	}

	authService struct {
		logger     *zap.SugaredLogger
		auth       authRepo
		validator  *validator.Validate
		jwtSecret  string
		passSecret string
	}

	bannerService struct {
		logger *zap.SugaredLogger
		banner bannerRepo
	}
)

func NewAuthService(logger *zap.SugaredLogger, jwtSecret, passSecret string, authRepo authRepo) *authService {
	return &authService{
		logger:     logger,
		auth:       authRepo,
		validator:  validator.New(),
		jwtSecret:  jwtSecret,
		passSecret: passSecret,
	}
}

func NewBannerService(logger *zap.SugaredLogger, bannerRepo bannerRepo) *bannerService {
	return &bannerService{
		logger: logger,
		banner: bannerRepo,
	}
}
