package cache

import (
	"testing"

	"github.com/gogapopp/trainee-assignment/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestCacheRepo(t *testing.T) {
	cache := New()
	banner := models.PostBannerRequest{
		TagIDs:    []int{1, 2},
		FeatureID: 1,
		IsActive:  true,
		Content: models.Content{
			Title: "Test Title",
			Text:  "Test Text",
			URL:   "Test URL",
		},
	}

	cache.SetUserBannerInCache(banner)

	tests := []struct {
		name      string
		tagId     int
		featureId int
		want      models.UserBannerResponse
		wantOk    bool
	}{
		{
			name:      "#1 ok",
			tagId:     1,
			featureId: 1,
			want: models.UserBannerResponse{
				IsActive: true,
				Content: models.Content{
					Title: "Test Title",
					Text:  "Test Text",
					URL:   "Test URL",
				},
			},
			wantOk: true,
		},
		{
			name:      "#2 ok",
			tagId:     2,
			featureId: 1,
			want: models.UserBannerResponse{
				IsActive: true,
				Content: models.Content{
					Title: "Test Title",
					Text:  "Test Text",
					URL:   "Test URL",
				},
			},
			wantOk: true,
		},
		{
			name:      "#2 false",
			tagId:     3,
			featureId: 1,
			want:      models.UserBannerResponse{},
			wantOk:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := cache.GetUserBannerFromCache(tt.tagId, tt.featureId)
			assert.Equal(t, tt.wantOk, ok)
			assert.Equal(t, tt.want, got)
		})
	}
}
