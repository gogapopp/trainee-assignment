package models

import "time"

type (
	// SignUpRequest the request body for the handler SignUp.
	SignUpRequest struct {
		UserID       int    `json:""`
		Username     string `json:"username" validate:"required"`
		PasswordHash string `json:"password" validate:"required"`
		Role         string `json:"role" validate:"required"`
	}
	// SignInRequest the request body for the handler SignIn.
	SignInRequest struct {
		Username     string `json:"username" validate:"required"`
		PasswordHash string `json:"password" validate:"required"`
	}
	// SignInResponse the response body for the handler SignIn.
	SignInResponse struct {
		Token string `json:"token"`
	}
	// Banner content.
	Content struct {
		Title string `json:"title" validate:"required"`
		Text  string `json:"text" validate:"required"`
		URL   string `json:"url" validate:"required"`
	}
	// PostBannerRequest the request body for the handler PostBanner.
	PostBannerRequest struct {
		BannerId  int     `json:""`
		TagIDs    []int   `json:"tag_ids" validate:"required"`
		FeatureID int     `json:"feature_id" validate:"required"`
		Content   Content `json:"content" validate:"required"`
		IsActive  bool    `json:"is_active"`
	}
	// PostBanner the response body for the handler PostBanner.
	PostBannerResponse struct {
		BannerID int `json:"banner_id"`
	}
	// UserBannerRequest describes the request parameters for the handler GetUserBanner.
	UserBannerRequest struct {
		TagId           int
		FeatureId       int
		UseLastRevision *bool
	}
	// UserBannerResponse the response body for the handler GetUserBanner.
	UserBannerResponse struct {
		Content  Content `json:""`
		IsActive bool    `json:""`
	}
	// GetBanner describes the request parameters for the handler BannersRequest.
	BannersRequest struct {
		FeatureId *int
		TagId     *int
		Limit     *int
		Offset    *int
	}
	// BannersResponse the response body for the handler GetBanner.
	BannersResponse struct {
		BannerId  int   `json:"banner_id"`
		TagIds    []int `json:"tag_ids"`
		FeatureId int   `json:"feature_id"`
		Content   Content
		IsActive  bool      `json:"is_active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	// PatchBanner the request body for the handler PatchBannerId.
	PatchBanner struct {
		TagIds    []int   `json:"tag_ids,omitempty"`
		FeatureId int     `json:"feature_id,omitempty"`
		Content   Content `json:"content,omitempty"`
		IsActive  *bool   `json:"is_active,omitempty"`
	}
)
