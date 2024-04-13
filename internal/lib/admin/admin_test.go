package admin

import (
	"context"
	"testing"

	"github.com/gogapopp/trainee-assignment/internal/handler/middlewares"
	"github.com/stretchr/testify/assert"
)

func TestIsAdmin(t *testing.T) {
	ctx := context.Background()
	ctxWithAdmin := context.WithValue(ctx, middlewares.UserRoleKey, "admin")
	ctxWithUser := context.WithValue(ctx, middlewares.UserRoleKey, "user")

	assert.True(t, IsAdmin(ctxWithAdmin))
	assert.False(t, IsAdmin(ctxWithUser))
}
