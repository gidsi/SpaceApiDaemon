package main

import "net/http"

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
}