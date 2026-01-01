package frontend_app

import (
	"crypto/rand"
	"math/big"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateRandomString generates a random string of the specified length.
func GenerateRandomString(length int) string {
	const characterSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	max, _ := rand.Int(rand.Reader, big.NewInt(int64(len(characterSet))))
	var bytes = make([]byte, length)
	for i := range bytes {
		bytes[i] = characterSet[max.Int64() % int64(len(characterSet))]
	}
	return string(bytes)
}

// GenerateJWT generates a JWT token with the specified claims.
func GenerateJWT(claims map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

// ValidateJWT checks if the provided JWT token is valid.
func ValidateJWT(tokenStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, nil
}

// GetNow gets the current UTC time.
func GetNow() time.Time {
	return time.Now().UTC()
}

// SplitString splits the input string into a slice of strings based on the specified delimiter.
func SplitString(input string, delimiter string) []string {
	return strings.Split(input, delimiter)
}