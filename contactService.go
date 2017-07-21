package main

import (
	"net/http"
	"spaceapi-spec/v013"
)

func changePhone(w http.ResponseWriter, r *http.Request) {
	value, spacedata := contactHelper(w, r)
	spacedata.Contact.Phone = value
	writeSpaceData(spacedata)
}

func changeSip(w http.ResponseWriter, r *http.Request) {
	value, spacedata := contactHelper(w, r)
	spacedata.Contact.Sip = value
	writeSpaceData(spacedata)
}

func changeIrc(w http.ResponseWriter, r *http.Request) {
	value, spacedata := contactHelper(w, r)
	spacedata.Contact.Irc = value
	writeSpaceData(spacedata)
}

func changeTwitter(w http.ResponseWriter, r *http.Request) {
	value, spacedata := contactHelper(w, r)
	spacedata.Contact.Twitter = value
	writeSpaceData(spacedata)
}

func changeFacebook(w http.ResponseWriter, r *http.Request) {
	value, spacedata := contactHelper(w, r)
	spacedata.Contact.Facebook = value
	writeSpaceData(spacedata)
}

func changeIdentica(w http.ResponseWriter, r *http.Request) {
	value, spacedata := contactHelper(w, r)
	spacedata.Contact.Identica = value
	writeSpaceData(spacedata)
}

func changeFoursquare(w http.ResponseWriter, r *http.Request) {
	value, spacedata := contactHelper(w, r)
	spacedata.Contact.Foursquare = value
	writeSpaceData(spacedata)
}

func changeEmail(w http.ResponseWriter, r *http.Request) {
	value, spacedata := contactHelper(w, r)
	spacedata.Contact.Email = value
	writeSpaceData(spacedata)
}

func changeMl(w http.ResponseWriter, r *http.Request) {
	value, spacedata := contactHelper(w, r)
	spacedata.Contact.Ml = value
	writeSpaceData(spacedata)
}

func changeJabber(w http.ResponseWriter, r *http.Request) {
	value, spacedata := contactHelper(w, r)
	spacedata.Contact.Jabber = value
	writeSpaceData(spacedata)
}

func changeIssueMail(w http.ResponseWriter, r *http.Request) {
	value, spacedata := contactHelper(w, r)
	spacedata.Contact.IssueMail = value
	writeSpaceData(spacedata)
}

func contactHelper(w http.ResponseWriter, r *http.Request) (string, spaceapi_spec.Root) {
	spacedata := readSpacedata()

	requestSpacedata := SingleValue{}
	createEntry(&requestSpacedata, w, r)

	if(spacedata.Contact == nil) {
		spacedata.Contact = &spaceapi_spec.Contact{}
	}

	return requestSpacedata.Value, spacedata
}