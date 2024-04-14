package admin

import (
	"context"

	"github.com/gogapopp/trainee-assignment/internal/handler/middlewares"
)

// IsAdmin проверяет из контекста роль admin для пользователя
func IsAdmin(ctx context.Context) bool {
	userRole := ctx.Value(middlewares.UserRoleKey)
	if r, ok := userRole.(string); ok {
		if r == "admin" {
			return true
		}
	}
	return false
}
