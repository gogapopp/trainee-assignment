package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gogapopp/trainee-assignment/internal/lib/jwt"
)

type (
	ctxKeyUserID int
	ctxKeyRole   string
)

const (
	UserIDKey   ctxKeyUserID = 0
	UserRoleKey ctxKeyRole   = ""
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

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserIDKey, fmt.Sprint(userID))
		ctx = context.WithValue(ctx, UserRoleKey, fmt.Sprint(role))

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func RolesMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if ok := isAdmin(ctx); !ok {
			http.Error(w, "you dont have premissions", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func isAdmin(ctx context.Context) bool {
	userRole := ctx.Value(UserRoleKey)
	if r, ok := userRole.(string); ok {
		if r == "admin" {
			return true
		}
	}
	return false
}
