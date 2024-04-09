package models

type (
	SignUpRequest struct {
		UserID       int    `json:""`
		Username     string `json:"username" validate:"required"`
		PasswordHash string `json:"password" validate:"required"`
		IsAdmin      bool   `json:"is_admin"`
	}

	SignInRequest struct {
		Username     string `json:"username" validate:"required"`
		PasswordHash string `json:"password" validate:"required"`
	}

	SignInResponse struct {
		Token string `json:"token"`
	}
)
