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
		SignIn(ctx context.Context, user models.SignInRequest) (string, error)
	}

	bannerRepo interface {
		SaveBanner(ctx context.Context, banner models.PostBannerRequest) (int, error)
		GetUserBanner(ctx context.Context, params models.UserBannerRequest) (models.UserBannerResponse, error)
		GetBanners(ctx context.Context, params models.BannersRequest) ([]models.BannersResponse, error)
		DeleteBanner(ctx context.Context, id int) error
		PatchBannerId(ctx context.Context, id int, banner models.PatchBanner) error
		DeleteBannerByFeatureId(ctx context.Context, featureId int) error
	}

	bannerCache interface {
		GetUserBannerFromCache(tagId, featureId int) (models.UserBannerResponse, bool)
		SetUserBannerInCache(banner models.PostBannerRequest)
	}
)

type (
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
		cache      bannerCache
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

func NewBannerService(logger *zap.SugaredLogger, bannerRepo bannerRepo, cache bannerCache) *bannerService {
	return &bannerService{
		logger:     logger,
		bannerRepo: bannerRepo,
		cache:      cache,
		validator:  validator.New(),
	}
}
