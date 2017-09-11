package main

import (
	"net/http"
	"github.com/gidsi/SpaceApiSpec/v013"
	"encoding/json"
	"strings"
	"sort"
)

var historyRoutes = routes{
	route{
		"",
		"GET",
		"/history/state",
		false,
		getHistoryState,
	},
	route{
		"Import History",
		"GET",
		"/importN39",
		true,
		importFromN39,
	},
}

type itemsSpaceAPIWithTimestamp struct {
	Items []spaceapi_spec.State `json:"items"`
}

type StateList []spaceapi_spec.State

func (s StateList) Len() int {
	return len(s)
}
func (s StateList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s StateList) Less(i, j int) bool {
	return s[i].Lastchange < s[j].Lastchange
}


func getHistoryState(w http.ResponseWriter, _ *http.Request) {
	spaceData, err := readSpaceData()

	stateSlice := StateList{}
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

		sort.Sort(stateSlice)

		returnJSON(w, itemsSpaceAPIWithTimestamp{stateSlice})
	}
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