package service

import (
	"context"
	"testing"
	"time"

	"github.com/gogapopp/trainee-assignment/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBannerRepo struct {
	mock.Mock
}

func (m *MockBannerRepo) SaveBanner(ctx context.Context, banner models.PostBannerRequest) (int, error) {
	args := m.Called(ctx, banner)
	return args.Int(0), args.Error(1)
}

func (m *MockBannerRepo) GetUserBanner(ctx context.Context, params models.UserBannerRequest) (models.UserBannerResponse, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(models.UserBannerResponse), args.Error(1)
}

func (m *MockBannerRepo) GetBanners(ctx context.Context, params models.BannersRequest) ([]models.BannersResponse, error) {
	args := m.Called(ctx, params)
	return args.Get(0).([]models.BannersResponse), args.Error(1)
}

func (m *MockBannerRepo) DeleteBanner(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockBannerRepo) PatchBannerId(ctx context.Context, id int, banner models.PatchBanner) error {
	args := m.Called(ctx, id, banner)
	return args.Error(0)
}

func (m *MockBannerRepo) DeleteBannerByFeatureId(ctx context.Context, featureId int) error {
	args := m.Called(ctx, featureId)
	return args.Error(0)
}

type MockCache struct {
	mock.Mock
}

func (m *MockCache) GetUserBannerFromCache(tagId, featureId int) (models.UserBannerResponse, bool) {
	args := m.Called(tagId, featureId)
	return args.Get(0).(models.UserBannerResponse), args.Bool(1)
}

func (m *MockCache) SetUserBannerInCache(banner models.PostBannerRequest) {
	m.Called(banner)
}

func TestBannerService_SaveBanner(t *testing.T) {
	bannerRepo := new(MockBannerRepo)
	cache := new(MockCache)
	s := NewBannerService(nil, bannerRepo, cache)

	tests := []struct {
		name   string
		banner models.PostBannerRequest
		mock   func()
		want   int
	}{
		{
			name: "#1 ok",
			banner: models.PostBannerRequest{
				TagIDs:    []int{1, 2, 3},
				FeatureID: 1,
				IsActive:  false,
				Content: models.Content{
					Title: "Test Title",
					Text:  "Test Text",
					URL:   "Test URL",
				},
			},
			mock: func() {
				bannerRepo.On("SaveBanner", mock.Anything, mock.Anything).Return(1, nil)
				cache.On("SetUserBannerInCache", mock.Anything)
			},
			want: 1,
		},
		{
			name: "#2 ok",
			banner: models.PostBannerRequest{
				TagIDs:    []int{1},
				FeatureID: 6,
				IsActive:  true,
				Content: models.Content{
					Title: "Test Title",
					Text:  "Test Text",
					URL:   "Test URL",
				},
			},
			mock: func() {
				bannerRepo.On("SaveBanner", mock.Anything, mock.Anything).Return(1, nil)
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := s.SaveBanner(context.Background(), tt.banner)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBannerService_GetUserBanner(t *testing.T) {
	bannerRepo := new(MockBannerRepo)
	cache := new(MockCache)
	s := NewBannerService(nil, bannerRepo, cache)

	tests := []struct {
		name   string
		params models.UserBannerRequest
		mock   func()
		want   models.UserBannerResponse
	}{
		{
			name:   "#1 ok",
			params: models.UserBannerRequest{TagId: 1, FeatureId: 1},
			mock: func() {
				banner := models.UserBannerResponse{IsActive: true}
				bannerRepo.On("GetUserBanner", mock.Anything, mock.Anything).Return(banner, nil)
				cache.On("GetUserBannerFromCache", mock.Anything, mock.Anything).Return(banner, true)
				cache.On("SetUserBannerInCache", mock.Anything)
			},
			want: models.UserBannerResponse{IsActive: true},
		},
		{
			name:   "#2 ok",
			params: models.UserBannerRequest{TagId: 1, FeatureId: 1},
			mock: func() {
				bannerRepo.On("GetUserBanner", mock.Anything, mock.Anything).Return(models.UserBannerResponse{}, nil)
				cache.On("GetUserBannerFromCache", mock.Anything, mock.Anything).Return(models.UserBannerResponse{}, false)
			},
			want: models.UserBannerResponse{IsActive: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := s.GetUserBanner(context.Background(), tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBannerService_GetBanners(t *testing.T) {
	bannerRepo := new(MockBannerRepo)
	cache := new(MockCache)
	s := NewBannerService(nil, bannerRepo, cache)

	tests := []struct {
		name   string
		params models.BannersRequest
		mock   func()
		want   []models.BannersResponse
	}{
		{
			name:   "#1 ok",
			params: models.BannersRequest{},
			mock: func() {
				banners := []models.BannersResponse{
					{
						BannerId:  0,
						TagIds:    []int{},
						FeatureId: 0,
						Content: models.Content{
							Title: "",
							Text:  "",
							URL:   "",
						},
						IsActive:  false,
						CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					},
				}
				bannerRepo.On("GetBanners", mock.Anything, mock.Anything).Return(banners, nil)
			},
			want: []models.BannersResponse{
				{
					BannerId:  0,
					TagIds:    []int{},
					FeatureId: 0,
					Content: models.Content{
						Title: "",
						Text:  "",
						URL:   "",
					},
					IsActive:  false,
					CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := s.GetBanners(context.Background(), tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBannerService_DeleteBanner(t *testing.T) {
	bannerRepo := new(MockBannerRepo)
	s := NewBannerService(nil, bannerRepo, nil)

	tests := []struct {
		name string
		id   int
		mock func()
	}{
		{
			name: "#1 ok",
			id:   1,
			mock: func() {
				bannerRepo.On("DeleteBanner", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "#2 ok",
			id:   2,
			mock: func() {
				bannerRepo.On("DeleteBanner", mock.Anything, mock.Anything).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := s.DeleteBanner(context.Background(), tt.id)
			assert.NoError(t, err)
		})
	}
}

func TestBannerService_PatchBannerId(t *testing.T) {
	bannerRepo := new(MockBannerRepo)
	s := NewBannerService(nil, bannerRepo, nil)

	tests := []struct {
		name   string
		id     int
		banner models.PatchBanner
		mock   func()
	}{
		{
			name:   "#1 ok",
			id:     1,
			banner: models.PatchBanner{},
			mock: func() {
				bannerRepo.On("PatchBannerId", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := s.PatchBannerId(context.Background(), tt.id, tt.banner)
			assert.NoError(t, err)
		})
	}
}

func TestBannerService_DeleteBannerByFeatureId(t *testing.T) {
	bannerRepo := new(MockBannerRepo)
	s := NewBannerService(nil, bannerRepo, nil)

	tests := []struct {
		name string
		id   int
		mock func()
	}{
		{
			name: "OK",
			id:   1,
			mock: func() {
				bannerRepo.On("DeleteBannerByFeatureId", mock.Anything, mock.Anything).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := s.DeleteBannerByFeatureId(context.Background(), tt.id)
			assert.NoError(t, err)
		})
	}
}
