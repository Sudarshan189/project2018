package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type MeterBillReading struct {
	RRNum string `json:"rr_num"` // Revenue Registration Number ( Unique ID )
	// PresentRead  string  `json:"present_read"`			// Present Load
	// PowerFactor  string  `json:"power_factor"`			// Power Factor

}

type Meter struct {
	RRNum        string  `json:"meter_data,omitempty"`
	CurrentValue float32 `json:"current_value,omitempty"`
	VoltageValue float32 `json:"voltage_value,omitempty"`
}

type MeterString struct {
	RRNum        string `json:"meter_data,omitempty"`
	CurrentValue string `json:"current_value,omitempty"`
	VoltageValue string `json:"voltage_value,omitempty"`
}

var meterstring MeterString
var meter Meter

func GetMeterData(w http.ResponseWriter, r *http.Request) { // optional

	// standard header set by browser...
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meter)
	Showlog(w, r)

}

func PostMeterData(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&meterstring)

	cur, err := strconv.ParseFloat(meterstring.CurrentValue, 32)
	if err != nil {
		panic(err.Error())
	}
	current := float32(cur)

	vol, err := strconv.ParseFloat(meterstring.VoltageValue, 32)
	if err != nil {
		panic(err.Error())
	}
	voltage := float32(vol)

	meter = Meter{meterstring.RRNum, current, voltage}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(meter)
	Showlog(w, r)
}
