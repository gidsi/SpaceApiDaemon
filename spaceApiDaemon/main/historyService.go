package main

import (
	"net/http"
	"github.com/spaceapi-community/go-spaceapi-spec/v13"
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
	Items []spaceapiStruct.State `json:"items"`
}

type StateList []spaceapiStruct.State

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
	stateSlice2 := StateList{}
	filteredStates := map[float64]spaceapiStruct.State{}

	if err == nil {
		for i := range spaceData {
			if spaceData[i].Data.State != nil && spaceData[i].Data.State.Lastchange != 0 {
				filteredStates[spaceData[i].Data.State.Lastchange] = *spaceData[i].Data.State
			}
		}

		for i := range filteredStates {
			stateSlice = append(stateSlice, filteredStates[i])
		}

		sort.Sort(stateSlice)
		lastState := !stateSlice[0].Open
		for i := range stateSlice {
            if stateSlice[i].Open != lastState {
                stateSlice2 = append(stateSlice2, stateSlice[i])
                lastState = stateSlice[i].Open
            }
        }

		returnJSON(w, itemsSpaceAPIWithTimestamp{stateSlice2})
	}
}