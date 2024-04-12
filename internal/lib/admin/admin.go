package admin

import (
	"context"

	"github.com/gogapopp/trainee-assignment/internal/handler/middlewares"
)

func IsAdmin(ctx context.Context) bool {
	userRole := ctx.Value(middlewares.UserRoleKey)
	if r, ok := userRole.(string); ok {
		if r == "admin" {
			return true
		}
	}
	return false
}
