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
	}

	APIHandler struct {
		logger *zap.SugaredLogger
		auth   authService
		banner bannerService
	}
)

func New(logger *zap.SugaredLogger, authService authService, bannerService bannerService) *APIHandler {
	return &APIHandler{
		logger: logger,
		auth:   authService,
		banner: bannerService,
	}
}
