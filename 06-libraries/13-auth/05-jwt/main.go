package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserInfo struct {
	Name string
	Role string
}

type CustomClaims struct {
	*jwt.StandardClaims
	TokenType 	string
	User 		UserInfo
}

func main() {
	key := []byte("SUPER_SECRET_KEY")

	tokenString := CreateToken(key)
	fmt.Println("tokenString:")
	fmt.Println(tokenString)
	fmt.Println()

	token := ParseToken(key, tokenString)
	claims := token.Claims.(*CustomClaims)
	fmt.Println("user:")
	fmt.Println(claims.User.Name)
}

func CreateToken(key []byte) string {
	c := &CustomClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		TokenType: "level1",
		User: UserInfo{
			Name: "gopher",
			Role: "admin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	tokenString, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	return tokenString
}

func ParseToken(key []byte, tokenString string) *jwt.Token {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		panic(err)
	}

	return token
}