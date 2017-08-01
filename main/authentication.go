package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"encoding/base64"
	"net/http"
)

type Token struct {
	Token string `json:"token"`
}

func addToken(w http.ResponseWriter, _ *http.Request) {
	token := Token{createToken()}

	writeToken(token.Token)

	ReturnJson(w, token)
}

func getToken(w http.ResponseWriter, _ *http.Request){
	tokenArray := readToken()

	ReturnJson(w, tokenArray)
}

func deleteToken(w http.ResponseWriter, r *http.Request) {
	token := Token{}
	createEntry(&token, w, r)

	removeToken(token.Token)

	w.WriteHeader(http.StatusNoContent)
}

func createToken() string {
	signingKey := []byte(config.SigningKey)

	claims := &jwt.StandardClaims{
		IssuedAt: time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(signingKey)
	encodedToken := base64.StdEncoding.EncodeToString([]byte(ss))

	return encodedToken
}