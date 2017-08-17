package main

import (
	"net/http"
	"github.com/gidsi/SpaceApiSpec/v013"
)

var feedRoutes = routes{
	route{
		"Blog Feed",
		"PUT",
		"/feeds/blog",
		true,
		changeBlogFeed,
	},
	route{
		"Wiki Feed",
		"PUT",
		"/feeds/wiki",
		true,
		changeWikiFeed,
	},
	route{
		"Calendar Feed",
		"PUT",
		"/feeds/calendar",
		true,
		changeCalendarFeed,
	},
	route{
		"Flickr Feed",
		"PUT",
		"/feeds/flickr",
		true,
		changeFlickrFeed,
	},
}

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
	spaceData, _ := readLastSpaceData()

	requestSpaceData := spaceapi_spec.Feed{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.Feeds == nil {
		spaceData.Feeds = &spaceapi_spec.Feeds{}
	}

	return spaceData, requestSpaceData
}