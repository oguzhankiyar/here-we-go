package auth

import "github.com/golang-jwt/jwt"

type AuthClaims struct {
	jwt.StandardClaims

	Name string `json:"name"`
}
