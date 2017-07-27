package main

import (
	"github.com/gidsi/SpaceApiSpec/v013"
	"net/http"
	"log"
	"strconv"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var config = ConfigFile{
	Port: "8080",
	SigningKey: "AllYourBase",
	MongoDbServer: "localhost",
	MongoDbDatabase: "spaceApi",
}

func main() {
	data, _ := ioutil.ReadFile("config.yaml")
	yaml.Unmarshal(data, &config)

	tokens := readToken()

	if(len(tokens) == 0) {
		token := createToken()
		writeToken(token)
		log.Printf("Initial Token created: %s", token)
	}

        router := SetupRouter(IndexRoutes)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(config.Port), router))
}

type base struct {
	Space string `json:"space"`
	Logo string `json:"logo"`
	Url string `json:"url"`
}

type StateShort struct {
	Open          bool `json:"open"`
	Lastchange    int `json:"lastchange,omitempty"`
	TriggerPerson string `json:"trigger_person,omitempty"`
	Message       string `json:"message,omitempty"`
}

type SingleValue struct {
	Value	string `json:"value"`
}

func changeBaseData(w http.ResponseWriter, r *http.Request) {
	inputData := base{}
	spacedata := readSpacedata()

	createEntry(&inputData, w, r)

	spacedata.Space = inputData.Space
	spacedata.Logo = inputData.Logo
	spacedata.Url = inputData.Url

	writeSpaceData(spacedata)
}

func changeState(w http.ResponseWriter, r *http.Request) {
	spacedata := readSpacedata()

	requestSpacedata := StateShort{}
	createEntry(&requestSpacedata, w, r)

	if(spacedata.State == nil) {
		spacedata.State = &spaceapi_spec.State{}
	}

	spacedata.State.Open = requestSpacedata.Open
	spacedata.State.Lastchange = requestSpacedata.Lastchange
	spacedata.State.Message = requestSpacedata.Message
	spacedata.State.TriggerPerson = requestSpacedata.TriggerPerson

	writeSpaceData(spacedata)
}

func changeLocation(w http.ResponseWriter, r *http.Request) {
	spacedata := readSpacedata()

	requestSpacedata := spaceapi_spec.Location{}
	createEntry(&requestSpacedata, w, r)
	spacedata.Location = &requestSpacedata

	writeSpaceData(spacedata)
}

