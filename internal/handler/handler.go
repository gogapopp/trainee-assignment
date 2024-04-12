package handler

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/gogapopp/trainee-assignment/internal/models"
	"go.uber.org/zap"
)

var valErr validator.ValidationErrors

type (
	authService interface {
		SignUp(ctx context.Context, user models.SignUpRequest) error
		SignIn(ctx context.Context, user models.SignInRequest) (string, error)
	}

	bannerService interface {
		SaveBanner(ctx context.Context, banner models.PostBannerRequest) (int, error)
	}

	APIHandler struct {
		logger        *zap.SugaredLogger
		authService   authService
		bannerService bannerService
	}
)

func New(logger *zap.SugaredLogger, authService authService, bannerService bannerService) *APIHandler {
	return &APIHandler{
		logger:        logger,
		authService:   authService,
		bannerService: bannerService,
	}
}
