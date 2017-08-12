package main

import (
	"net/http"
	"github.com/gidsi/SpaceApiSpec/v013"
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
