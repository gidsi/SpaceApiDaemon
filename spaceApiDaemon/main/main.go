package main

import (
	"net/http"
	"log"
	"strconv"
)

func main() {
	initConfig()
	createInitToken()

	log.Print(readToken())

	routes := []routes{
		staticRoutes,
		indexRoutes,
		issueReportChannelRoutes,
		locationRoutes,
		authenticationRoutes,
		historyRoutes,
		importRoutes,
		baseDataRoutes,
		contactRoutes,
		feedRoutes,
		sensorsRoutes,
		stateRoutes,
	}

	var ApplicationRoutes []route
	for _, r := range routes {
		ApplicationRoutes = append(ApplicationRoutes, r...)
	}

	router := setupRouter(ApplicationRoutes)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(config.APIPort), router))
}

type singleValue struct {
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