package main

import (
	"net/http"
	"github.com/gidsi/SpaceApiSpec/v013"
	"github.com/gorilla/mux"
)

func addDoorLocked(w http.ResponseWriter, r *http.Request) {
	spaceData := readSpaceData()

	requestSpaceData := spaceapi_spec.DoorLocked{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.Sensors == nil {
		spaceData.Sensors = &spaceapi_spec.Sensors{}
	}

	spaceData.Sensors.DoorLocked = append(spaceData.Sensors.DoorLocked, requestSpaceData)

	writeSpaceData(spaceData)
}

func changeDoorLocked(w http.ResponseWriter, r *http.Request) {
	spaceData := readSpaceData()

	requestSpaceData := spaceapi_spec.DoorLocked{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.Sensors == nil {
		spaceData.Sensors = &spaceapi_spec.Sensors{}
	}

	vars := mux.Vars(r)
	for i := 0; i < len(spaceData.Sensors.DoorLocked); i++ {
		if spaceData.Sensors.DoorLocked[i].Location == vars["location"] {
			spaceData.Sensors.DoorLocked[i] = requestSpaceData
			break
		}
	}

	writeSpaceData(spaceData)
}


func removeDoorLocked(w http.ResponseWriter, r *http.Request) {
	spaceData := readSpaceData()

	requestSpaceData := spaceapi_spec.DoorLocked{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.Sensors == nil {
		spaceData.Sensors = &spaceapi_spec.Sensors{}
	}

	vars := mux.Vars(r)
	for i := 0; i < len(spaceData.Sensors.DoorLocked); i++ {
		if spaceData.Sensors.DoorLocked[i].Location == vars["location"] {
			spaceData.Sensors.DoorLocked =
				append(
					spaceData.Sensors.DoorLocked[:i],
					spaceData.Sensors.DoorLocked[i+1:]...
				)
			break
		}
	}

	writeSpaceData(spaceData)
}

func addTemperature(w http.ResponseWriter, r *http.Request) {
	spaceData := readSpaceData()

	requestSpaceData := spaceapi_spec.Temperature{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.Sensors == nil {
		spaceData.Sensors = &spaceapi_spec.Sensors{}
	}

	spaceData.Sensors.Temperature = append(spaceData.Sensors.Temperature, requestSpaceData)

	writeSpaceData(spaceData)
}


func changeTemperature(w http.ResponseWriter, r *http.Request) {
	spaceData := readSpaceData()

	requestSpaceData := spaceapi_spec.Temperature{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.Sensors == nil {
		spaceData.Sensors = &spaceapi_spec.Sensors{}
	}

	vars := mux.Vars(r)
	for i := 0; i < len(spaceData.Sensors.Temperature); i++ {
		if spaceData.Sensors.Temperature[i].Location == vars["location"] {
			spaceData.Sensors.Temperature[i] = requestSpaceData
			break
		}
	}

	writeSpaceData(spaceData)
}


func removeTemperature(w http.ResponseWriter, r *http.Request) {
	spaceData := readSpaceData()

	requestSpaceData := spaceapi_spec.Temperature{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.Sensors == nil {
		spaceData.Sensors = &spaceapi_spec.Sensors{}
	}

	vars := mux.Vars(r)
	for i := 0; i < len(spaceData.Sensors.Temperature); i++ {
		if spaceData.Sensors.Temperature[i].Location == vars["location"] {
			spaceData.Sensors.Temperature =
				append(
					spaceData.Sensors.Temperature[:i],
					spaceData.Sensors.Temperature[i+1:]...
				)
			break
		}
	}

	writeSpaceData(spaceData)
}