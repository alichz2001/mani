package main

import (
	jtoken "github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateTokenForUser(user *User) (string, error) {
	claims := jtoken.MapClaims{
		"ID":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
