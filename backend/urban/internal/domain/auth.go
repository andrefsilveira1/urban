package domain

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	user string
	jwt.StandardClaims
}
