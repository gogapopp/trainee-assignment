package models

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

	content struct {
		Title string `json:"title" validate:"required"`
		Text  string `json:"text" validate:"required"`
		URL   string `json:"url" validate:"required"`
	}

	PostBannerRequest struct {
		TagIDs    []int   `json:"tag_ids" validate:"required"`
		FeatureID int     `json:"feature_id" validate:"required"`
		Content   content `json:"content" validate:"required"`
		IsActive  bool    `json:"is_active"`
	}

	PostBannerResponse struct {
		BannerID int `json:"banner_id"`
	}
)
