package main

import (
	"net/http"
	"github.com/gidsi/SpaceApiSpec/v013"
	"github.com/gorilla/mux"
)

var contactRoutes = routes{
	route{
		"Phone",
		"PUT",
		"/contact/phone",
		true,
		changePhone,
	},
	route{
		"Sip",
		"PUT",
		"/contact/sip",
		true,
		changeSip,
	},
	route{
		"Irc",
		"PUT",
		"/contact/irc",
		true,
		changeIrc,
	},
	route{
		"Twitter",
		"PUT",
		"/contact/twitter",
		true,
		changeTwitter,
	},
	route{
		"Facebook",
		"PUT",
		"/contact/facebook",
		true,
		changeFacebook,
	},
	route{
		"Identica",
		"PUT",
		"/contact/identica",
		true,
		changeIdentica,
	},
	route{
		"Foursquare",
		"PUT",
		"/contact/foursquare",
		true,
		changeFoursquare,
	},
	route{
		"Email",
		"PUT",
		"/contact/email",
		true,
		changeEmail,
	},
	route{
		"Ml",
		"PUT",
		"/contact/ml",
		true,
		changeMl,
	},
	route{
		"Jabber",
		"PUT",
		"/contact/jabber",
		true,
		changeJabber,
	},
	route{
		"IssueMail",
		"PUT",
		"/contact/issuemail",
		true,
		changeIssueMail,
	},
	route{
		"Add Keymaster",
		"POST",
		"/contact/keymasters",
		true,
		addKeymaster,
	},
	route{
		"Remove Keymaster",
		"DELETE",
		"/contact/keymasters/{name}",
		true,
		removeKeymaster,
	},
}

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

func addKeymaster(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()
	if spaceData.Contact == nil {
		spaceData.Contact = &spaceapi_spec.Contact{}
	}

	requestKeymaster := spaceapi_spec.Keymaster{}
	createEntry(&requestKeymaster, w, r)

	spaceData.Contact.Keymasters = append(spaceData.Contact.Keymasters, requestKeymaster)

	writeSpaceData(spaceData)
}

func removeKeymaster(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()
	if spaceData.Contact == nil {
		spaceData.Contact = &spaceapi_spec.Contact{}
	}

	vars := mux.Vars(r)
	for i := 0; i < len(spaceData.Contact.Keymasters); i++ {
		if spaceData.Contact.Keymasters[i].Name == vars["name"] {
			spaceData.Contact.Keymasters =
				append(
					spaceData.Contact.Keymasters[:i],
					spaceData.Contact.Keymasters[i+1:]...
				)
			break
		}
	}

	writeSpaceData(spaceData)
	w.WriteHeader(http.StatusNoContent)
}

func contactHelper(w http.ResponseWriter, r *http.Request) (string, spaceapi_spec.Root) {
	spaceData, _ := readLastSpaceData()

	requestSpaceData := singleValue{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.Contact == nil {
		spaceData.Contact = &spaceapi_spec.Contact{}
	}

	return requestSpaceData.Value, spaceData
}