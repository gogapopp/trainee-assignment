package models

import "time"

type (
	SignUpRequest struct {
		UserID       int    `json:""`
		Username     string `json:"username" validate:"required"`
		PasswordHash string `json:"password" validate:"required"`
		Role         string `json:"role" validate:"required"`
	}

	SignInRequest struct {
		Username     string `json:"username" validate:"required"`
		PasswordHash string `json:"password" validate:"required"`
	}

	SignInResponse struct {
		Token string `json:"token"`
	}

	Content struct {
		Title string `json:"title" validate:"required"`
		Text  string `json:"text" validate:"required"`
		URL   string `json:"url" validate:"required"`
	}

	PostBannerRequest struct {
		BannerId  int     `json:""`
		TagIDs    []int   `json:"tag_ids" validate:"required"`
		FeatureID int     `json:"feature_id" validate:"required"`
		Content   Content `json:"content" validate:"required"`
		IsActive  bool    `json:"is_active"`
	}

	PostBannerResponse struct {
		BannerID int `json:"banner_id"`
	}

	UserBannerRequest struct {
		TagId           int
		FeatureId       int
		UseLastRevision *bool
	}

	UserBannerResponse struct {
		Content  Content `json:""`
		IsActive bool    `json:""`
	}

	BannersRequest struct {
		FeatureId *int
		TagId     *int
		Limit     *int
		Offset    *int
	}

	BannersResponse struct {
		BannerId  int   `json:"banner_id"`
		TagIds    []int `json:"tag_ids"`
		FeatureId int   `json:"feature_id"`
		Content   Content
		IsActive  bool      `json:"is_active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	PatchBanner struct {
		TagIds    []int   `json:"tag_ids,omitempty"`
		FeatureId int     `json:"feature_id,omitempty"`
		Content   Content `json:"content,omitempty"`
		IsActive  *bool   `json:"is_active,omitempty"`
	}
)
