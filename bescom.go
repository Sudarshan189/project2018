package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

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
	db.Where(&KWHUpdate{RRNum: rrnum}).First(&meterstring)
	fmt.Println(meterstring)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meterstring)
	Showlog(w, r)

}
