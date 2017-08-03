package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func addIssueReportChannel(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readSpaceData()

	inputData := SingleValue{}
	createEntry(&inputData, w, r)

	spaceData.IssueReportChannels = append(spaceData.IssueReportChannels, inputData.Value)

	writeSpaceData(spaceData)
}

func removeIssueReportChannel(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readSpaceData()

	vars := mux.Vars(r)
	for i := 0; i < len(spaceData.IssueReportChannels); i++ {
		if spaceData.IssueReportChannels[i] == vars["name"] {
			spaceData.IssueReportChannels =
				append(
					spaceData.IssueReportChannels[:i],
					spaceData.IssueReportChannels[i+1:]...
				)
			break
		}
	}

	w.WriteHeader(http.StatusNoContent)
}