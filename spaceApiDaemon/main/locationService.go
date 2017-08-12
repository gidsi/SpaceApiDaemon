package main

import (
	"net/http"
	"github.com/gidsi/SpaceApiSpec/v013"
)

func changeLocation(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()

	requestSpaceData := spaceapi_spec.Location{}
	createEntry(&requestSpaceData, w, r)
	spaceData.Location = &requestSpaceData

	writeSpaceData(spaceData)
}
