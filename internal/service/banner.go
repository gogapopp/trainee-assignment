package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/gogapopp/trainee-assignment/internal/handler/middlewares"
	"github.com/gogapopp/trainee-assignment/internal/models"
)

var (
	ErrBannerUnactive   = errors.New("banner is unactive")
	ErrBannersNotExist  = errors.New("banners dont exist")
	ErrNoFieldsToUpdate = errors.New("no field to update")
)

func (b *bannerService) SaveBanner(ctx context.Context, banner models.PostBannerRequest) (int, error) {
	const op = "service.banner.SaveBanner"
	err := b.validator.Struct(banner)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	bannerId, err := b.bannerRepo.SaveBanner(ctx, banner)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	banner.BannerId = bannerId
	b.cache.SetUserBannerInCache(banner)
	return bannerId, nil
}

func (b *bannerService) GetUserBanner(ctx context.Context, params models.UserBannerRequest) (models.UserBannerResponse, error) {
	const op = "service.banner.GetUserBanner"
	banner, ok := b.cache.GetUserBannerFromCache(params.TagId, params.FeatureId)
	if !ok || (params.UseLastRevision != nil && *params.UseLastRevision) {
		banner, err := b.bannerRepo.GetUserBanner(ctx, params)
		if err != nil {
			return models.UserBannerResponse{}, fmt.Errorf("%s: %w", op, err)
		}
		err = isActive(ctx, banner.IsActive)
		if err != nil {
			return models.UserBannerResponse{}, fmt.Errorf("%s: %w", op, err)
		}
		b.cache.SetUserBannerInCache(models.PostBannerRequest{
			TagIDs:    []int{params.TagId},
			FeatureID: params.FeatureId,
			IsActive:  banner.IsActive,
			Content: models.Content{
				Title: banner.Content.Title,
				Text:  banner.Content.Text,
				URL:   banner.Content.URL,
			},
		})
		return banner, nil
	}
	err := isActive(ctx, banner.IsActive)
	if err != nil {
		return models.UserBannerResponse{}, fmt.Errorf("%s: %w", op, err)
	}
	return banner, nil
}

func isActive(ctx context.Context, isActive bool) error {
	if ok := middlewares.IsAdmin(ctx); ok {
		return nil
	}
	if !isActive {
		return ErrBannerUnactive
	}
	return nil
}

func (b *bannerService) GetBanners(ctx context.Context, params models.BannersRequest) ([]models.BannersResponse, error) {
	const op = "service.banner.GetBanners"
	banners, err := b.bannerRepo.GetBanners(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if len(banners) < 1 {
		return nil, fmt.Errorf("%s: %w", op, ErrBannersNotExist)
	}
	return banners, nil
}

func (b *bannerService) DeleteBanner(ctx context.Context, id int) error {
	return b.bannerRepo.DeleteBanner(ctx, id)
}

func (b *bannerService) PatchBannerId(ctx context.Context, id int, banner models.PatchBanner) error {
	const op = "service.banner.PatchBannerId"
	err := b.bannerRepo.PatchBannerId(ctx, id, banner)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (b *bannerService) DeleteBannerByFeatureId(ctx context.Context, featureId int) error {
	return b.bannerRepo.DeleteBannerByFeatureId(ctx, featureId)
}
