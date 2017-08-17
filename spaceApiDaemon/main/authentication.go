package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
)

var authenticationRoutes = routes{
	route{
		"list token",
		"GET",
		"/token",
		true,
		getToken,
	},
	route{
		"Generate new token",
		"POST",
		"/token",
		true,
		addToken,
	},
	route{
		"Remove token",
		"DELETE",
		"/token",
		true,
		deleteToken,
	},
}

type Token struct {
	Token string `json:"token"`
}

func addToken(w http.ResponseWriter, _ *http.Request) {
	token := Token{createToken()}

	writeToken(token.Token)

	returnJSON(w, token)
}

func getToken(w http.ResponseWriter, _ *http.Request){
	tokenArray := readToken()

	returnJSON(w, tokenArray)
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

	return ss
}