package main

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

func main() {
	jwt.New(jwt.SigningMethodES256)
}
