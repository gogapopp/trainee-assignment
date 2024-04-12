package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/gogapopp/trainee-assignment/internal/models"
)

func (s *storage) SaveBanner(ctx context.Context, banner models.PostBannerRequest) (int, error) {
	const (
		op    = "postgres.banner.SaveBanner"
		query = "INSERT INTO banners (tag_ids, feature_id, banner_data, is_active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING banner_id;"
	)
	var bannerId int
	row := s.db.QueryRow(ctx, query, banner.TagIDs, banner.FeatureID, banner.Content, banner.IsActive, time.Now(), time.Now())
	err := row.Scan(&bannerId)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return bannerId, nil
}
