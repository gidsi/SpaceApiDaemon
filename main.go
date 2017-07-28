package main

import (
	"github.com/gidsi/SpaceApiSpec/v013"
	"net/http"
	"log"
	"strconv"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"spaceapi-spec/v013"
)

var config = ConfigFile{
	Port: 8080,
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
	spaceData := readSpaceData()

	createEntry(&inputData, w, r)

	spaceData.Space = inputData.Space
	spaceData.Logo = inputData.Logo
	spaceData.Url = inputData.Url

	writeSpaceData(spaceData)
}

func changeState(w http.ResponseWriter, r *http.Request) {
	spaceData := readSpaceData()

	requestSpaceData := StateShort{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.State == nil {
		spaceData.State = &spaceapi_spec.State{}
	}

	spaceData.State.Open = requestSpaceData.Open
	spaceData.State.Lastchange = requestSpaceData.Lastchange
	spaceData.State.Message = requestSpaceData.Message
	spaceData.State.TriggerPerson = requestSpaceData.TriggerPerson

	writeSpaceData(spaceData)
}

func changeLocation(w http.ResponseWriter, r *http.Request) {
	spaceData := readSpaceData()

	requestSpaceData := spaceapi_spec.Location{}
	createEntry(&requestSpaceData, w, r)
	spaceData.Location = &requestSpaceData

	writeSpaceData(spaceData)
}

