package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"strings"
	"time"
)

func setupRouter(routes routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	methods := make(map[string][]string)
	for _, route := range routes {
		setupRouteForRouter(router, route)
		methods[route.Pattern] = append(methods[route.Pattern], route.Method)
	}

	for _, routeMethods := range methods {
		router.
			Methods("OPTIONS").
			Name("Options Handler").
			Handler(optionsHandler(append(routeMethods, "OPTIONS")))
	}


	router.NotFoundHandler = http.HandlerFunc(notFound)

	return router
}

func setupRouteForRouter(router *mux.Router, route route) {
	router.
	Methods(route.Method).
	Path(route.Pattern).
	Name(route.Name).
	Handler(setupRoute(route))
}

// TODO: clean that mess
func setupRoute(route route) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		setCORSHeader([]string{ route.Method }, w, r)
		if route.AuthNeeded {
			if checkToken(r.Header.Get("Authorization")) ||
				checkToken(r.URL.Query().Get("Authorization")) {
				route.Handler.ServeHTTP(w, r)
			} else {
				log.Println("Provided Token not found!")
				w.WriteHeader(http.StatusUnauthorized)
			}
		} else {
			route.Handler.ServeHTTP(w, r)
		}

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			route.Name,
			time.Since(start),
		)
	})
}

func notFound(w http.ResponseWriter, _ *http.Request) {
	log.Print("Route not found!")
	w.WriteHeader(http.StatusNotFound)
}

func optionsHandler(methods []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setCORSHeader(methods, w, r)
		w.WriteHeader(200)
	})
}

func setCORSHeader(methods []string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", checkAllowedOrigin(r))
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods[:],", "))
}

// TODO: config.AllowedOrigins should be a regular expression
func checkAllowedOrigin(request *http.Request) string {
	for _, b := range config.AllowedOrigins {
        if b == request.Header.Get("Origin") {
            return request.Header.Get("Origin")
        }
    }

	return ""
}