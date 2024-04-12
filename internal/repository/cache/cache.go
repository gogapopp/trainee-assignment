package cache

import (
	"fmt"
	"time"

	"github.com/gogapopp/trainee-assignment/internal/models"
	"github.com/patrickmn/go-cache"
)

type cacheRepo struct {
	cache *cache.Cache
}

func New() *cacheRepo {
	cache := cache.New(5*time.Minute, 10*time.Minute)
	return &cacheRepo{cache: cache}
}

func (c *cacheRepo) GetUserBannerFromCache(tagId, featureId int) (models.UserBannerResponse, bool) {
	cacheKey := fmt.Sprintf("banner_%d_%d", tagId, featureId)
	if banner, ok := c.cache.Get(cacheKey); ok {
		if banner, ok := banner.(models.UserBannerResponse); ok {
			return banner, true
		}
		return models.UserBannerResponse{}, false
	}
	return models.UserBannerResponse{}, false
}

func (c *cacheRepo) SetUserBannerInCache(banner models.PostBannerRequest) {
	for k := range banner.TagIDs {
		cacheKey := fmt.Sprintf("banner_%d_%d", banner.TagIDs[k], banner.FeatureID)
		var bannerResp models.UserBannerResponse
		bannerResp.IsActive = banner.IsActive
		bannerResp.Content.Title = banner.Content.Title
		bannerResp.Content.Text = banner.Content.Text
		bannerResp.Content.URL = banner.Content.URL
		c.cache.Set(cacheKey, bannerResp, time.Minute*5)
	}
}
