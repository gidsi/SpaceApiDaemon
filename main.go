package main

import (
	"net/http"
	"log"
	"strconv"
)

func main() {
	initConfig()
	createInitToken()

	router := SetupRouter(IndexRoutes)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(config.Port), router))
}

type SingleValue struct {
	Value	string `json:"value"`
}

func createInitToken() {
	tokens := readToken()
	if len(tokens) == 0 {
		token := createToken()
		writeToken(token)
		log.Printf("Initial Token created: %s", token)
	}
}