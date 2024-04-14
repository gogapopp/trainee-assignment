package models

import "time"

type (
	// SignUpRequest тело запроса для хендлера SignUp
	SignUpRequest struct {
		UserID       int    `json:""`
		Username     string `json:"username" validate:"required"`
		PasswordHash string `json:"password" validate:"required"`
		Role         string `json:"role" validate:"required"`
	}
	// SignInRequest тело запроса для хендлера SignIn
	SignInRequest struct {
		Username     string `json:"username" validate:"required"`
		PasswordHash string `json:"password" validate:"required"`
	}
	// SignInResponse тело ответа для хендлера SignIn
	SignInResponse struct {
		Token string `json:"token"`
	}

	Content struct {
		Title string `json:"title" validate:"required"`
		Text  string `json:"text" validate:"required"`
		URL   string `json:"url" validate:"required"`
	}
	// PostBannerRequest описывает тело запроса для хендлера PostBanner
	PostBannerRequest struct {
		BannerId  int     `json:""`
		TagIDs    []int   `json:"tag_ids" validate:"required"`
		FeatureID int     `json:"feature_id" validate:"required"`
		Content   Content `json:"content" validate:"required"`
		IsActive  bool    `json:"is_active"`
	}
	// PostBanner описывает тело ответа для хендлера PostBanner
	PostBannerResponse struct {
		BannerID int `json:"banner_id"`
	}
	// UserBannerRequest описывает параметры запроса для хендлера GetUserBanner
	UserBannerRequest struct {
		TagId           int
		FeatureId       int
		UseLastRevision *bool
	}
	// UserBannerResponse описывает тело ответа для хендлера GetUserBanner
	UserBannerResponse struct {
		Content  Content `json:""`
		IsActive bool    `json:""`
	}
	// GetBanner тело параметры запроса хендлера BannersRequest
	BannersRequest struct {
		FeatureId *int
		TagId     *int
		Limit     *int
		Offset    *int
	}
	// BannersResponse тело ответа для хендлера GetBanner
	BannersResponse struct {
		BannerId  int   `json:"banner_id"`
		TagIds    []int `json:"tag_ids"`
		FeatureId int   `json:"feature_id"`
		Content   Content
		IsActive  bool      `json:"is_active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	// PatchBanner описывает тело запроса для PatchBannerId
	PatchBanner struct {
		TagIds    []int   `json:"tag_ids,omitempty"`
		FeatureId int     `json:"feature_id,omitempty"`
		Content   Content `json:"content,omitempty"`
		IsActive  *bool   `json:"is_active,omitempty"`
	}
)
