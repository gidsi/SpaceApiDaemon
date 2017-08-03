package main

import (
	"net/http"
	"github.com/gidsi/SpaceApiSpec/v013"
)

func changeBlogFeed(w http.ResponseWriter, r *http.Request) {
	spaceData, feedData := feedsHelper(w, r)

	spaceData.Feeds.Blog = &feedData

	writeSpaceData(spaceData)
}

func changeWikiFeed(w http.ResponseWriter, r *http.Request) {
	spaceData, feedData := feedsHelper(w, r)

	spaceData.Feeds.Blog = &feedData

	writeSpaceData(spaceData)
}

func changeCalendarFeed(w http.ResponseWriter, r *http.Request) {
	spaceData, feedData := feedsHelper(w, r)

	spaceData.Feeds.Blog = &feedData

	writeSpaceData(spaceData)
}

func changeFlickrFeed(w http.ResponseWriter, r *http.Request) {
	spaceData, feedData := feedsHelper(w, r)

	spaceData.Feeds.Blog = &feedData

	writeSpaceData(spaceData)
}

func feedsHelper(w http.ResponseWriter, r *http.Request) (spaceapi_spec.Root, spaceapi_spec.Feed) {
	spaceData, _ := readSpaceData()

	requestSpaceData := spaceapi_spec.Feed{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.Feeds == nil {
		spaceData.Feeds = &spaceapi_spec.Feeds{}
	}

	return spaceData, requestSpaceData
}