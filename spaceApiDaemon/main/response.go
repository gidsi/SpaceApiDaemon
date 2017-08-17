package main

import (
	"net/http"
	"encoding/json"
)

var indexRoutes = routes{
	route{
		"Index",
		"GET",
		"/",
		false,
		index,
	},
}

func index(w http.ResponseWriter, _ *http.Request) {
	spaceData, err := readLastSpaceData()

	if err == nil {
		returnJSON(w, spaceData)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func returnJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
}

func createEntry(i interface{}, w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(i); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(i); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			panic(err)
		}
	}
}