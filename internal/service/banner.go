package service

import (
	"context"
	"fmt"

	"github.com/gogapopp/trainee-assignment/internal/models"
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
	return bannerId, nil
}
