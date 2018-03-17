package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
)

const statusSuccess  = 1
const statusFailed = 2

type MeterString struct {
	MeterNum string `json:"meter_data,omitempty"`
	CurrentValue float32 `json:"current_value,omitempty"`
	VoltageValue float32 `json:"voltage_value,omitempty"`
}

type Meter struct {
	MeterNum string `json:"meter_data,omitempty"`
	CurrentValue string `json:"current_value,omitempty"`
	VoltageValue string `json:"voltage_value,omitempty"`
}

type Confirm struct {
	Status int `json:"status,omitempty"`
}

// var meters[] Meter  // for all return
var meter Meter  // for single
var meterstring MeterString  // for single
var confirm Confirm


func GetMeterData(w http.ResponseWriter, r *http.Request) {
	log.Printf("/data GET")
	json.NewEncoder(w).Encode(meterstring)
}

func PostMeterData(w http.ResponseWriter, r *http.Request) {
	log.Printf("/meterdata POST")
	//params := mux.Vars(r)

	_=json.NewDecoder(r.Body).Decode(&meter)

	cur, err := strconv.ParseFloat(meter.CurrentValue, 32)
	if err != nil {
		// do something sensible
	}
	current := float32(cur)
	vol, err := strconv.ParseFloat(meter.VoltageValue, 32)
	if err != nil {
		// do something sensible
	}
	voltage := float32(vol)
	meterstring = MeterString{meter.MeterNum, current, voltage}

	json.NewEncoder(w).Encode(meterstring)
}

//func PostMeterData(w http.ResponseWriter, r *http.Request) {
//	log.Printf("/meterdata POST")
//	//params := mux.Vars(r)
//
//	_=json.NewDecoder(r.Body).Decode(&d)
//	a = Now{d}
//	confirm = Confirm{statusSuccess}
//	json.NewEncoder(w).Encode(a)
//}


func main() {
	router := mux.NewRouter()
	// meters = append(meters, Meter{"KA1256", 0.25, 220})
	// meters = append(meters, Meter{"KA1258", 0.86, 220})

	router.HandleFunc("/data", GetMeterData).Methods("GET")

	router.HandleFunc("/meterdata", PostMeterData).Methods("POST")

	log.Printf("Serving on :8080, Go to localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}





