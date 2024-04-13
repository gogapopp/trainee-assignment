package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndParseJWTToken(t *testing.T) {
	jwtSecret := "test_secret"
	role := "admin"
	username := "test_user"
	password := "test_password"

	token, err := GenerateJWTToken(jwtSecret, role, username, password)
	assert.NoError(t, err)

	assert.NotEmpty(t, token)

	parsedRole, err := ParseJWTToken(jwtSecret, token)
	assert.NoError(t, err)

	assert.Equal(t, role, parsedRole)

	roleUser := "user"
	usernameUser := "test_user"
	passwordUser := "test_password_user"

	tokenUser, err := GenerateJWTToken(jwtSecret, roleUser, usernameUser, passwordUser)
	assert.NoError(t, err)

	assert.NotEmpty(t, tokenUser)

	parsedRoleUser, err := ParseJWTToken(jwtSecret, tokenUser)
	assert.NoError(t, err)

	assert.Equal(t, roleUser, parsedRoleUser)
}
