package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

type KWH struct {
	RRNum     string  `json:"rr_num"` 		// Revenue Registration Number ( Unique ID )
	Current   string  `json:"current"`
	Voltage   string  `json:"voltage"`
	KWH       float64 `json:"kwh"` 			// Consumption ( Present - Previous )
	UpdatedAt string  `json:"updated_at"`
}


var me string
var meterstring KWHUpdate

func GetMeterData(w http.ResponseWriter, r *http.Request) {
	meterstring = KWHUpdate{}
	db, err := gorm.Open("mysql", "session:session@/project2018")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	vars := mux.Vars(r)
	rrnum := vars["rr_num"]

	// db.Where(&KWHUpdate{RRNum: rrnum}).First(&meterstring)
	db.Where(&KWHUpdate{RRNum:rrnum}).First(&meterstring)

	fmt.Println(meterstring)

	//meterstring.UpdatedAt = meterstring.UpdatedAt
	// standard header set by browser...
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meterstring)
	Showlog(w, r)

}
