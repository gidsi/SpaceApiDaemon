package main

import (
	"net/http"
	"github.com/spaceapi-community/go-spaceapi-spec/v13"
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
	spaceData, _ := readLastSpaceData()
	feedsHelper(&spaceData)

	requestSpaceData := spaceapiStruct.Blog{}
	createEntry(&requestSpaceData, w, r)

	spaceData.Feeds.Blog = &requestSpaceData

	writeSpaceData(spaceData)
}

func changeWikiFeed(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()
	feedsHelper(&spaceData)

	requestSpaceData := spaceapiStruct.Wiki{}
	createEntry(&requestSpaceData, w, r)

	spaceData.Feeds.Wiki = &requestSpaceData

	writeSpaceData(spaceData)
}

func changeCalendarFeed(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()
	feedsHelper(&spaceData)

	requestSpaceData := spaceapiStruct.Calendar{}
	createEntry(&requestSpaceData, w, r)

	spaceData.Feeds.Calendar = &requestSpaceData

	writeSpaceData(spaceData)
}

func changeFlickrFeed(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()
	feedsHelper(&spaceData)

	requestSpaceData := spaceapiStruct.Flickr{}
	createEntry(&requestSpaceData, w, r)

	spaceData.Feeds.Flickr = &requestSpaceData

	writeSpaceData(spaceData)
}

func feedsHelper(spaceData *spaceapiStruct.SpaceAPI013) {
	if spaceData.Feeds == nil {
		spaceData.Feeds = &spaceapiStruct.Feeds{}
	}
}