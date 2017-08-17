package main

import (
	"net/http"
	"github.com/gidsi/SpaceApiSpec/v013"
)

var stateRoutes = routes{
	route{
		"State",
		"PUT",
		"/state",
		true,
		changeState,
	},
}

type stateShort struct {
	Open          bool `json:"open"`
	Lastchange    int `json:"lastchange,omitempty"`
	TriggerPerson string `json:"trigger_person,omitempty"`
	Message       string `json:"message,omitempty"`
}

func changeState(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()

	requestSpaceData := stateShort{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.State == nil {
		spaceData.State = &spaceapi_spec.State{}
	}

	spaceData.State.Open = requestSpaceData.Open
	spaceData.State.Lastchange = requestSpaceData.Lastchange
	spaceData.State.Message = requestSpaceData.Message
	spaceData.State.TriggerPerson = requestSpaceData.TriggerPerson

	writeSpaceData(spaceData)
}