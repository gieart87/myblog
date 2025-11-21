package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// getJWTSecret retrieves JWT secret from environment variable
func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// Fallback to a default value for development (should be set in production)
		secret = "your-secret-key"
	}
	return []byte(secret)
}

// JWTClaim represents the JWT claims structure
type JWTClaim struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a new JWT token for a given user ID
func GenerateJWT(userID string) (string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &JWTClaim{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTSecret())
}

// ValidateJWT validates a JWT token and returns the claims if valid
func ValidateJWT(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return getJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
