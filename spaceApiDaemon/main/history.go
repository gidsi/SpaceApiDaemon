package main

import (
	"net/http"
	"github.com/gidsi/SpaceApiSpec/v013"
	"encoding/json"
	"strings"
)

type ItemsSpaceApiWithTimestamp struct {
	Items []spaceapi_spec.State `json:"items"`
}

func getHistoryState(w http.ResponseWriter, _ *http.Request) {
	spaceData, err := readSpaceData()

	stateSlice := []spaceapi_spec.State{}
	filteredStates := map[int]spaceapi_spec.State{}

	if err == nil {
		for i := range spaceData {
			if spaceData[i].Data.State != nil {
				filteredStates[spaceData[i].Data.State.Lastchange] = *spaceData[i].Data.State
			}
		}

		for i := range filteredStates {
			stateSlice = append(stateSlice, filteredStates[i])
		}

		ReturnJson(w, ItemsSpaceApiWithTimestamp{stateSlice})
	}
}

type N39State struct {
	Open bool `json:"open"`
	Lastchange int64 `json:"lastchange"`
}

type N39Item struct {
	Id    string `json:"id"`
	Key   int64 `json:"key"`
	Value N39State `json:"value"`
}

type N39Items struct {
	TotalRows	int `json:"total_rows"`
	Offset		int `json:"offset"`
	Rows		[]N39Item `json:"rows"`
}

func importFromN39(w http.ResponseWriter, _ *http.Request) {
	spaceData, _ := readLastSpaceData()
	foo := N39Items{}

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