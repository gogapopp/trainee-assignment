package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/gogapopp/trainee-assignment/internal/models"
	"github.com/gogapopp/trainee-assignment/internal/repository"
	"github.com/jackc/pgx/v5"
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

func (s *storage) GetUserBanner(ctx context.Context, params models.UserBannerRequest) (models.UserBannerResponse, error) {
	const (
		op    = "postgres.banner.GetUserBanner"
		query = "SELECT banner_data, is_active FROM banners WHERE $1=ANY(tag_ids) AND feature_id=$2;"
	)
	var banner models.UserBannerResponse
	row := s.db.QueryRow(ctx, query, params.TagId, params.FeatureId)
	err := row.Scan(&banner.Content, &banner.IsActive)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.UserBannerResponse{}, fmt.Errorf("%s: %w", op, repository.ErrBannerNotExist)
		}
		return models.UserBannerResponse{}, fmt.Errorf("%s: %w", op, err)
	}
	return banner, nil
}

func (s *storage) GetBanners(ctx context.Context, params models.BannersRequest) ([]models.BannersResponse, error) {
	const (
		op    = "postgres.banner.GetBanners"
		query = "SELECT * FROM banners WHERE ($1::integer IS NULL OR feature_id=$1) AND ($2::integer IS NULL OR $2=ANY(tag_ids)) ORDER BY created_at DESC LIMIT $3 OFFSET $4;"
	)
	var banners []models.BannersResponse
	rows, err := s.db.Query(ctx, query, params.FeatureId, params.TagId, params.Limit, params.Offset)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var banner models.BannersResponse
		err := rows.Scan(&banner.BannerId, &banner.TagIds, &banner.FeatureId, &banner.Content, &banner.IsActive, &banner.CreatedAt, &banner.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		banners = append(banners, banner)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return banners, nil
}

func (s *storage) DeleteBanner(ctx context.Context, id int) error {
	const (
		op    = "postgres.banner.DeleteBanner"
		query = "DELETE FROM banners WHERE banner_id=$1;"
	)
	pgt, err := s.db.Exec(ctx, query, id)
	if pgt.RowsAffected() < 1 {
		return fmt.Errorf("%s: %w", op, repository.ErrBannerNotExist)
	}
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *storage) PatchBannerId(ctx context.Context, id int, banner models.PatchBanner) error {
	const op = "postgres.banner. PatchBannerId"
	query := "UPDATE banners SET"
	args := []interface{}{}
	i := 1

	if banner.FeatureId != 0 {
		query += fmt.Sprintf(" feature_id = $%d,", i)
		args = append(args, banner.FeatureId)
		i++
	}
	if len(banner.TagIds) > 0 {
		query += fmt.Sprintf(" tag_ids = $%d,", i)
		args = append(args, banner.TagIds)
		i++
	}
	if banner.Content.Title != "" || banner.Content.Text != "" || banner.Content.URL != "" {
		query += fmt.Sprintf(" banner_data = $%d,", i)
		args = append(args, banner.Content)
		i++
	}
	if banner.IsActive {
		query += fmt.Sprintf(" is_active = $%d,", i)
		args = append(args, banner.IsActive)
		i++
	}

	if i == 1 {
		return fmt.Errorf("%s: %w", op, repository.ErrNoFieldsToUpdate)
	}

	query = query[:len(query)-1] + fmt.Sprintf(" WHERE banner_id = $%d", i)
	args = append(args, id)

	pgt, err := s.db.Exec(ctx, query, args...)
	if pgt.RowsAffected() < 1 {
		return fmt.Errorf("%s: %w", op, repository.ErrBannerNotExist)
	}
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
