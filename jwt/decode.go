package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func Decode(s string) (*jwt.Token, error) {
	// res := new(DecodedToken)
	token, err := jwt.Parse(s, nil)
	if token == nil {
		return token, fmt.Errorf("malformed token: %w", err)
	}
	return token, nil

}
