package main

import (
	"net/http"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	AuthNeeded bool
	Handler    http.HandlerFunc
}

type Routes []Route

var IndexRoutes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		false,
		Index,
	},
	Route{
		"Base Data",
		"PUT",
		"/base",
		true,
		changeBaseData,
	},
	Route{
		"Location",
		"PUT",
		"/location",
		true,
		changeLocation,
	},
	Route{
		"State",
		"PUT",
		"/state",
		true,
		changeState,
	},
	Route{
		"Phone",
		"PUT",
		"/contact/phone",
		true,
		changePhone,
	},
	Route{
		"Sip",
		"PUT",
		"/contact/sip",
		true,
		changeSip,
	},
	Route{
		"Irc",
		"PUT",
		"/contact/irc",
		true,
		changeIrc,
	},
	Route{
		"Twitter",
		"PUT",
		"/contact/twitter",
		true,
		changeTwitter,
	},
	Route{
		"Facebook",
		"PUT",
		"/contact/facebook",
		true,
		changeFacebook,
	},
	Route{
		"Identica",
		"PUT",
		"/contact/identica",
		true,
		changeIdentica,
	},
	Route{
		"Foursquare",
		"PUT",
		"/contact/foursquare",
		true,
		changeFoursquare,
	},
	Route{
		"Email",
		"PUT",
		"/contact/email",
		true,
		changeEmail,
	},
	Route{
		"Ml",
		"PUT",
		"/contact/ml",
		true,
		changeMl,
	},
	Route{
		"Jabber",
		"PUT",
		"/contact/jabber",
		true,
		changeJabber,
	},
	Route{
		"IssueMail",
		"PUT",
		"/contact/issuemail",
		true,
		changeIssueMail,
	},
	Route{
		"Add Keymaster",
		"POST",
		"/contact/keymasters",
		true,
		addKeymaster,
	},
	Route{
		"Remove Keymaster",
		"DELETE",
		"/contact/keymasters/{name}",
		true,
		removeKeymaster,
	},
	Route{
		"Blog Feed",
		"PUT",
		"/feeds/blog",
		true,
		changeBlogFeed,
	},
	Route{
		"Wiki Feed",
		"PUT",
		"/feeds/wiki",
		true,
		changeWikiFeed,
	},
	Route{
		"Calendar Feed",
		"PUT",
		"/feeds/calendar",
		true,
		changeCalendarFeed,
	},
	Route{
		"Flickr Feed",
		"PUT",
		"/feeds/flickr",
		true,
		changeFlickrFeed,
	},
	Route{
		"Add issue report channel",
		"POST",
		"/issue_report_channels",
		true,
		addIssueReportChannel,
	},
	Route{
		"Remove issue report channel",
		"DELETE",
		"/issue_report_channels/{name}",
		true,
		removeIssueReportChannel,
	},
	Route{
		"Add temperature sensor",
		"POST",
		"/sensors/temperature",
		true,
		addTemperature,
	},
	Route{
		"change temperature sensor",
		"PUT",
		"/sensors/temperature/{location}",
		true,
		changeTemperature,
	},
	Route{
		"Remove temperature sensor",
		"DELETE",
		"/sensors/temperature/{location}",
		true,
		removeTemperature,
	},
	Route{
		"Add door locked sensor",
		"POST",
		"/sensors/door_locked",
		true,
		addDoorLocked,
	},
	Route{
		"change door locked sensor",
		"PUT",
		"/sensors/door_locked/{location}",
		true,
		changeDoorLocked,
	},
	Route{
		"Remove Door locked sensor",
		"DELETE",
		"/sensors/door_locked/{location}",
		true,
		removeDoorLocked,
	},
	Route{
		"Add humidity sensor",
		"POST",
		"/sensors/humidity",
		true,
		addHumidity,
	},
	Route{
		"change humidity sensor",
		"PUT",
		"/sensors/humidity/{location}",
		true,
		changeHumidity,
	},
	Route{
		"Remove humidity sensor",
		"DELETE",
		"/sensors/humidity/{location}",
		true,
		removeHumidity,
	},
	Route{
		"list token",
		"GET",
		"/token",
		true,
		getToken,
	},
	Route{
		"Generate new token",
		"POST",
		"/token",
		true,
		addToken,
	},
	Route{
		"Remove token",
		"DELETE",
		"/token",
		true,
		deleteToken,
	},
}