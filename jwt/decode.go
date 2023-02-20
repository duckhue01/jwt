package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hokaccha/go-prettyjson"
)

func Decode(s string) ([2]string, error) {
	res := [2]string{}
	token, err := jwt.Parse(s, nil)
	if token == nil {
		return res, fmt.Errorf("malformed token: %w", err)
	}

	header, err := prettyjson.Marshal(token.Header)
	if err != nil {
		return res, fmt.Errorf("fail to encode header: %w", err)
	}
	res[0] = string(header)

	body, err := prettyjson.Marshal(token.Claims)
	if err != nil {
		return res, fmt.Errorf("fail to encode body: %w", err)
	}
	res[1] = string(body)

	return res, nil
}
