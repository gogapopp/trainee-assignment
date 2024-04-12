package service

import (
	"context"
	"errors"

	"github.com/go-playground/validator"
	"github.com/gogapopp/trainee-assignment/internal/models"
	"go.uber.org/zap"
)

var ErrUndefinedRole = errors.New("undefined role")

type (
	authRepo interface {
		SignUp(ctx context.Context, user models.SignUpRequest) error
		SignIn(ctx context.Context, user models.SignInRequest) (int, string, error)
	}

	bannerRepo interface {
		SaveBanner(ctx context.Context, banner models.PostBannerRequest) (int, error)
	}

	authService struct {
		logger     *zap.SugaredLogger
		authRepo   authRepo
		validator  *validator.Validate
		jwtSecret  string
		passSecret string
	}

	bannerService struct {
		logger     *zap.SugaredLogger
		bannerRepo bannerRepo
		validator  *validator.Validate
	}
)

func NewAuthService(jwtSecret, passSecret string, logger *zap.SugaredLogger, authRepo authRepo) *authService {
	return &authService{
		logger:     logger,
		authRepo:   authRepo,
		validator:  validator.New(),
		jwtSecret:  jwtSecret,
		passSecret: passSecret,
	}
}

func NewBannerService(logger *zap.SugaredLogger, bannerRepo bannerRepo) *bannerService {
	return &bannerService{
		logger:     logger,
		bannerRepo: bannerRepo,
		validator:  validator.New(),
	}
}
