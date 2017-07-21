package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func SetupRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		setupRoute(router, route)
	}

	router.
	Methods("OPTIONS").
	Name("Options Handler").
	Handler(http.HandlerFunc(optionsHandler))

	router.NotFoundHandler = http.HandlerFunc(notFound)

	return router
}

func setupRoute(router *mux.Router, route Route) {
	var handler http.Handler
	if(route.AuthNeeded) {
		handler = checkSecurity(route.Handler)
	} else {
		handler = Logger(route.Handler, route.Name)
	}

	router.
	Methods(route.Method).
	Path(route.Pattern).
	Name(route.Name).
	Handler(handler)

	router.
	Methods("OPTIONS").
	Name("Options Handler").
	Handler(http.HandlerFunc(optionsHandler))
}

func checkSecurity(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if(checkToken(r.Header.Get("auth"))) {
			inner.ServeHTTP(w, r)
		} else {
			log.Println("Provided Token not found!")
			w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		}
	})
}

func notFound(w http.ResponseWriter, r *http.Request) {
	log.Print("Route not found!")
	w.WriteHeader(http.StatusNotFound)
}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.WriteHeader(200)
}