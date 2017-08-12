package main

import "net/http"

type base struct {
	Space string `json:"space"`
	Logo string `json:"logo"`
	Url string `json:"url"`
}

func changeBaseData(w http.ResponseWriter, r *http.Request) {
	inputData := base{}
	spaceData, _ := readLastSpaceData()

	createEntry(&inputData, w, r)

	spaceData.Api = "0.13"
	spaceData.Space = inputData.Space
	spaceData.Logo = inputData.Logo
	spaceData.Url = inputData.Url

	writeSpaceData(spaceData)
}