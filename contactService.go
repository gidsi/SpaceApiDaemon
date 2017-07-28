package main

import (
	"net/http"
	"github.com/gidsi/SpaceApiSpec/v013"
)

func changePhone(w http.ResponseWriter, r *http.Request) {
	value, spaceData := contactHelper(w, r)
	spaceData.Contact.Phone = value
	writeSpaceData(spaceData)
}

func changeSip(w http.ResponseWriter, r *http.Request) {
	value, spaceData := contactHelper(w, r)
	spaceData.Contact.Sip = value
	writeSpaceData(spaceData)
}

func changeIrc(w http.ResponseWriter, r *http.Request) {
	value, spaceData := contactHelper(w, r)
	spaceData.Contact.Irc = value
	writeSpaceData(spaceData)
}

func changeTwitter(w http.ResponseWriter, r *http.Request) {
	value, spaceData := contactHelper(w, r)
	spaceData.Contact.Twitter = value
	writeSpaceData(spaceData)
}

func changeFacebook(w http.ResponseWriter, r *http.Request) {
	value, spaceData := contactHelper(w, r)
	spaceData.Contact.Facebook = value
	writeSpaceData(spaceData)
}

func changeIdentica(w http.ResponseWriter, r *http.Request) {
	value, spaceData := contactHelper(w, r)
	spaceData.Contact.Identica = value
	writeSpaceData(spaceData)
}

func changeFoursquare(w http.ResponseWriter, r *http.Request) {
	value, spaceData := contactHelper(w, r)
	spaceData.Contact.Foursquare = value
	writeSpaceData(spaceData)
}

func changeEmail(w http.ResponseWriter, r *http.Request) {
	value, spaceData := contactHelper(w, r)
	spaceData.Contact.Email = value
	writeSpaceData(spaceData)
}

func changeMl(w http.ResponseWriter, r *http.Request) {
	value, spaceData := contactHelper(w, r)
	spaceData.Contact.Ml = value
	writeSpaceData(spaceData)
}

func changeJabber(w http.ResponseWriter, r *http.Request) {
	value, spaceData := contactHelper(w, r)
	spaceData.Contact.Jabber = value
	writeSpaceData(spaceData)
}

func changeIssueMail(w http.ResponseWriter, r *http.Request) {
	value, spaceData := contactHelper(w, r)
	spaceData.Contact.IssueMail = value
	writeSpaceData(spaceData)
}

func contactHelper(w http.ResponseWriter, r *http.Request) (string, spaceapi_spec.Root) {
	spaceData := readSpaceData()

	requestSpaceData := SingleValue{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.Contact == nil {
		spaceData.Contact = &spaceapi_spec.Contact{}
	}

	return requestSpaceData.Value, spaceData
}