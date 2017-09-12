package main

import (
	"net/http"
	"github.com/gidsi/SpaceApiSpec/v013"
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