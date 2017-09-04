package main

import (
	"net/http"
	"github.com/gidsi/SpaceApiSpec/v013"
	"time"
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

	if spaceData.State.Open != requestSpaceData.Open && requestSpaceData.Lastchange == 0 {
		spaceData.State.Lastchange = int(time.Now().Unix())
	} else {
		spaceData.State.Lastchange = requestSpaceData.Lastchange
	}

	spaceData.State.Open = requestSpaceData.Open
	spaceData.State.Message = requestSpaceData.Message
	spaceData.State.TriggerPerson = requestSpaceData.TriggerPerson

	writeSpaceData(spaceData)
}