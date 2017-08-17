package main

import (
	"net/http"
	"io/ioutil"
	"log"
)

type route struct {
	Name       string
	Method     string
	Pattern    string
	AuthNeeded bool
	Handler    http.HandlerFunc
}

type routes []route

var staticRoutes = routes{
	route{
		"Swagger.json route",
		"GET",
		"/swagger.json",
		false,
		func (w http.ResponseWriter, r *http.Request) {
			b, err := ioutil.ReadFile("swagger.json")
			if err != nil {
				log.Fatal(err)
			}
			w.Write(b)
		},
	},
}