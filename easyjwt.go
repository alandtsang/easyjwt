// Package easyjwt is used to generate jwt token and parse jwt token,
// use HS256 to sign the payload, support custom secret and expiration time.
package easyjwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	defaultSecret     = "secret"
	defaultExpireTime = time.Hour * 1
)

var (
	// ErrPayloadEmpty means the payload passed in is empty.
	ErrPayloadEmpty = fmt.Errorf("Payload is empty")

	// ErrExpiredToken means the token has expired.
	ErrExpiredToken = fmt.Errorf("Token was expired")

	// ErrSignatureInvalid means that the signature of the token is invalid.
	ErrSignatureInvalid = fmt.Errorf("Signature validation failed")

	// ErrUnknown means unknown error type.
	ErrUnknown = fmt.Errorf("Unknown error")
)

// GenerateToken generates jwt token with data.
func GenerateToken(payload map[string]interface{}) (string, error) {
	return GenerateCustomToken(payload, defaultSecret, defaultExpireTime)
}

// GenerateCustomToken generates jwt token according to data, custom secret and expire time.
func GenerateCustomToken(payload map[string]interface{}, secret string, expire time.Duration) (string, error) {
	if len(payload) == 0 {
		return "", ErrPayloadEmpty
	}

	if secret == "" {
		secret = defaultSecret
	}

	if expire == 0 {
		expire = defaultExpireTime
	}

	// Set claims
	claims := jwt.MapClaims{}
	for k, v := range payload {
		claims[k] = v
	}
	claims["exp"] = time.Now().Add(expire).Unix()

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	return token.SignedString([]byte(secret))
}

// ParseToken parses jwt token and return data.
func ParseToken(tokenStr string) (interface{}, error) {
	return ParseCustomToken(tokenStr, defaultSecret)
}

// ParseCustomToken parses jwt token with custom secret and return data.
func ParseCustomToken(tokenStr, secret string) (interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrSignatureInvalid
		}
		return []byte(secret), nil
	})
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			return nil, ErrExpiredToken
		case jwt.ValidationErrorSignatureInvalid:
			return nil, ErrSignatureInvalid
		default:
			return nil, ErrUnknown
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		result := make(map[string]interface{}, len(claims))
		for k, v := range claims {
			result[k] = v
		}
		return result, nil
	}
	return nil, ErrUnknown
}
