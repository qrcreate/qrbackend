package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)


var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateJWT membuat token JWT untuk pengguna
func GenerateJWT(userID string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}
