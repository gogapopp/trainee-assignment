package middlewares

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gogapopp/trainee-assignment/internal/lib/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	os.Setenv("JWT_SECRET_KEY", "secret")
	token, _ := jwt.GenerateJWTToken(os.Getenv("JWT_SECRET_KEY"), "admin", "testuser", "testpassword")

	tests := []struct {
		name       string
		token      string
		wantStatus int
	}{
		{
			name:       "Valid token",
			token:      token,
			wantStatus: http.StatusOK,
		},
		{
			name:       "Invalid token",
			token:      "invalidtoken",
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "No token",
			token:      "",
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/", nil)
			request.Header.Set("token", tt.token)
			response := httptest.NewRecorder()

			handler := AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			handler.ServeHTTP(response, request)

			assert.Equal(t, tt.wantStatus, response.Code)
		})
	}
}
