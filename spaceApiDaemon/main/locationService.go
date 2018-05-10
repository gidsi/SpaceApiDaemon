package main

import (
	"net/http"
	"github.com/spaceapi-community/go-spaceapi-spec/v13"
)

var locationRoutes = routes{
	route{
		"Location",
		"PUT",
		"/location",
		true,
		changeLocation,
	},
}

func changeLocation(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()

	requestSpaceData := spaceapiStruct.Location{}
	createEntry(&requestSpaceData, w, r)
	spaceData.Location = &requestSpaceData

	writeSpaceData(spaceData)
}
