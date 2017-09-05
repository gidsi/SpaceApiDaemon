package main

import (
	"net/http"
	"github.com/gidsi/SpaceApiSpec/v013"
	"github.com/gorilla/mux"
)

var sensorsRoutes = routes{
	route{
		"Add temperature sensor",
		"POST",
		"/sensors/temperature",
		true,
		addTemperature,
	},
	route{
		"change temperature sensor",
		"PUT",
		"/sensors/temperature/{location}",
		true,
		changeTemperature,
	},
	route{
		"Remove temperature sensor",
		"DELETE",
		"/sensors/temperature/{location}",
		true,
		removeTemperature,
	},
	route{
		"Add door locked sensor",
		"POST",
		"/sensors/door_locked",
		true,
		addDoorLocked,
	},
	route{
		"change door locked sensor",
		"PUT",
		"/sensors/door_locked/{location}",
		true,
		changeDoorLocked,
	},
	route{
		"Remove Door locked sensor",
		"DELETE",
		"/sensors/door_locked/{location}",
		true,
		removeDoorLocked,
	},
	route{
		"Add humidity sensor",
		"POST",
		"/sensors/humidity",
		true,
		addHumidity,
	},
	route{
		"change humidity sensor",
		"PUT",
		"/sensors/humidity/{location}",
		true,
		changeHumidity,
	},
	route{
		"Remove humidity sensor",
		"DELETE",
		"/sensors/humidity/{location}",
		true,
		removeHumidity,
	},
}

func addDoorLocked(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()

	requestSpaceData := spaceapi_spec.DoorLocked{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.Sensors == nil {
		spaceData.Sensors = &spaceapi_spec.Sensors{}
	}

	spaceData.Sensors.DoorLocked = append(spaceData.Sensors.DoorLocked, requestSpaceData)

	writeSpaceData(spaceData)
}

func changeDoorLocked(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()
	vars := mux.Vars(r)

	requestSpaceData := spaceapi_spec.DoorLocked{}
	createEntry(&requestSpaceData, w, r)
	requestSpaceData.Location = vars["location"]

	if spaceData.Sensors == nil {
		spaceData.Sensors = &spaceapi_spec.Sensors{}
	}

	for i := 0; i < len(spaceData.Sensors.DoorLocked); i++ {
		if spaceData.Sensors.DoorLocked[i].Location == vars["location"] {
			spaceData.Sensors.DoorLocked[i] = requestSpaceData
			break
		}
	}

	writeSpaceData(spaceData)
}


func removeDoorLocked(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()

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
	spaceData, _ := readLastSpaceData()

	requestSpaceData := spaceapi_spec.Temperature{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.Sensors == nil {
		spaceData.Sensors = &spaceapi_spec.Sensors{}
	}

	spaceData.Sensors.Temperature = append(spaceData.Sensors.Temperature, requestSpaceData)

	writeSpaceData(spaceData)
}


func changeTemperature(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()
	vars := mux.Vars(r)

	requestSpaceData := spaceapi_spec.Temperature{}
	createEntry(&requestSpaceData, w, r)
	requestSpaceData.Location = vars["location"]

	if spaceData.Sensors == nil {
		spaceData.Sensors = &spaceapi_spec.Sensors{}
	}

	for i := 0; i < len(spaceData.Sensors.Temperature); i++ {
		if spaceData.Sensors.Temperature[i].Location == vars["location"] {
			spaceData.Sensors.Temperature[i] = requestSpaceData
			break
		}
	}

	writeSpaceData(spaceData)
}


func removeTemperature(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()

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

func addHumidity(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()

	requestSpaceData := spaceapi_spec.Humidity{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.Sensors == nil {
		spaceData.Sensors = &spaceapi_spec.Sensors{}
	}

	spaceData.Sensors.Humidity = append(spaceData.Sensors.Humidity, requestSpaceData)

	writeSpaceData(spaceData)
}

func changeHumidity(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()
	vars := mux.Vars(r)

	requestSpaceData := spaceapi_spec.Humidity{}
	createEntry(&requestSpaceData, w, r)
	requestSpaceData.Location = vars["location"]

	if spaceData.Sensors == nil {
		spaceData.Sensors = &spaceapi_spec.Sensors{}
	}

	for i := 0; i < len(spaceData.Sensors.Humidity); i++ {
		if spaceData.Sensors.Humidity[i].Location == vars["location"] {
			spaceData.Sensors.Humidity[i] = requestSpaceData
			break
		}
	}

	writeSpaceData(spaceData)
}

func removeHumidity(w http.ResponseWriter, r *http.Request) {
	spaceData, _ := readLastSpaceData()

	requestSpaceData := spaceapi_spec.Humidity{}
	createEntry(&requestSpaceData, w, r)

	if spaceData.Sensors == nil {
		spaceData.Sensors = &spaceapi_spec.Sensors{}
	}

	vars := mux.Vars(r)
	for i := 0; i < len(spaceData.Sensors.Humidity); i++ {
		if spaceData.Sensors.Humidity[i].Location == vars["location"] {
			spaceData.Sensors.Humidity =
				append(
					spaceData.Sensors.Humidity[:i],
					spaceData.Sensors.Humidity[i+1:]...
				)
			break
		}
	}

	writeSpaceData(spaceData)
}