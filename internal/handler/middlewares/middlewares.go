package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gogapopp/trainee-assignment/internal/lib/jwt"
)

type (
	CtxKeyUserID   int
	CtxKeyUserRole bool
)

const (
	UserIDKey   CtxKeyUserID   = 0
	UserRoleKey CtxKeyUserRole = false
)

func AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("token")
		if authHeader == "" {
			http.Error(w, "authorization header not found", http.StatusUnauthorized)
			return
		}

		userID, role, err := jwt.ParseJWTToken(os.Getenv("JWT_SECRET_KEY"), authHeader)
		if err != nil {
			http.Error(w, "invalid authorization token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, fmt.Sprint(userID))
		ctx = context.WithValue(ctx, UserRoleKey, fmt.Sprint(role))

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
