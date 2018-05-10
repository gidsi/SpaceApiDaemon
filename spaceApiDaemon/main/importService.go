package main

import (
	"net/http"
	"encoding/json"
	"strings"
	"github.com/spaceapi-community/go-spaceapi-spec/v13"
)

var importRoutes = routes{
	route{
		"Import History",
		"GET",
		"/importN39History",
		true,
		importFromN39,
	},
	route{
		"Import SpaceApi status",
		"PUT",
		"/importJsonFile",
		true,
		importSpaceApiStatus,
	},
}

type n39State struct {
	Open bool `json:"open"`
	Lastchange int64 `json:"lastchange"`
}

type n39Item struct {
	ID    string `json:"id"`
	Key   int64 `json:"key"`
	Value n39State `json:"value"`
}

type n39Items struct {
	TotalRows	int `json:"total_rows"`
	Offset		int `json:"offset"`
	Rows		[]n39Item `json:"rows"`
}

func importFromN39(_ http.ResponseWriter, _ *http.Request) {
	spaceData, _ := readLastSpaceData()
	foo := n39Items{}

	r, err := http.Get("http://spaceapi-stats.n39.eu/" + strings.ToLower(spaceData.Space) + "/_design/space/_view/all")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&foo)

	for i := range foo.Rows {
		writeImportedData(foo.Rows[i])
	}
}

func importSpaceApiStatus(w http.ResponseWriter, r *http.Request) {
	spaceApiStatus := spaceapiStruct.SpaceAPI013{}
	createEntry(&spaceApiStatus, w, r)
	writeSpaceData(spaceApiStatus)
}