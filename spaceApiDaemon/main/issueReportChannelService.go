package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

var issueReportChannelRoutes = routes{
	route{
		"Add issue report channel",
		"POST",
		"/issue_report_channels",
		true,
		addIssueReportChannel,
	},
	route{
		"Remove issue report channel",
		"DELETE",
		"/issue_report_channels/{name}",
		true,
		removeIssueReportChannel,
	},
}

func addIssueReportChannel(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()

	inputData := singleValue{}
	createEntry(&inputData, w, r)

	spaceData.IssueReportChannels = append(spaceData.IssueReportChannels, inputData.Value)

	writeSpaceData(spaceData)
}

func removeIssueReportChannel(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()

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