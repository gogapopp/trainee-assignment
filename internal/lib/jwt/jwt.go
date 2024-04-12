package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var errUnknownClaimsType = errors.New("unknown claims type")

const TOKEN_EXP = time.Minute * 10

type tokenClaims struct {
	Role string
	jwt.RegisteredClaims
}

func GenerateJWTToken(jwtSecret string, role, username, password string) (string, error) {
	claims := tokenClaims{
		role,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_EXP)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func ParseJWTToken(jwtSecret, userJWTToken string) (string, error) {
	token, err := jwt.ParseWithClaims(userJWTToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte([]byte(jwtSecret)), nil
	})
	if err != nil {
		return "", err
	} else if claims, ok := token.Claims.(*tokenClaims); ok {
		return claims.Role, nil
	}
	return "", errUnknownClaimsType
}
