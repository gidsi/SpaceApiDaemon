package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"encoding/base64"
)

type Token struct {
	Token string `json:"token"`
}

func createToken() string {
	mySigningKey := []byte("AllYourBase")

	claims := &jwt.StandardClaims{
		IssuedAt: time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	encodedToken := base64.StdEncoding.EncodeToString([]byte(ss))

	return encodedToken
}