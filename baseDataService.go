package main

import "net/http"

type base struct {
	Space string `json:"space"`
	Logo string `json:"logo"`
	Url string `json:"url"`
}

func changeBaseData(w http.ResponseWriter, r *http.Request) {
	inputData := base{}
	spaceData := readSpaceData()

	createEntry(&inputData, w, r)

	spaceData.Space = inputData.Space
	spaceData.Logo = inputData.Logo
	spaceData.Url = inputData.Url

	writeSpaceData(spaceData)
}