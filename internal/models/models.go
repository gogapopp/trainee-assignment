package models

type (
	SignUpRequest struct {
		UserID       int    `json:""`
		Username     string `json:"username"`
		PasswordHash string `json:"password"`
		IsAdmin      bool   `json:"is_admin"`
	}

	SignInRequest struct {
		Username     string `json:"username"`
		PasswordHash string `json:"password"`
	}

	SignInResponse struct {
		Token string `json:"token"`
	}
)
