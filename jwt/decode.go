package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hokaccha/go-prettyjson"
)

type DecodedToken struct {
	Header string
	Body   string
}

func Decode(s string) (*DecodedToken, error) {
	res := new(DecodedToken)
	token, err := jwt.Parse(s, nil)
	if token == nil {
		return res, fmt.Errorf("malformed token: %w", err)
	}

	header, err := prettyjson.Marshal(token.Header)
	if err != nil {
		return res, fmt.Errorf("can't encode header: %w", err)
	}
	res.Header = string(header)

	body, err := prettyjson.Marshal(token.Claims)
	if err != nil {
		return res, fmt.Errorf("can't encode body: %w", err)
	}
	res.Body = string(body)

	return res, nil
}
