package main

import (
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"strconv"
)

type BilledData struct {
	First uint32
	Restall uint32
}

func CreateBillInfo(w http.ResponseWriter, r * http.Request) {
	var data KWHUpdate
	db, err := gorm.Open("mysql", "akshay:deepika019@/project2018")
	if err!= nil {
		panic(err.Error())
	}
	defer db.Close()
	Showlog(w,r)
	vars := mux.Vars(r)
	rrnum := vars["rr_num"]

	db.Where(&KWHUpdate{RRNum:rrnum}).First(&data)
	kwdata,err:= strconv.ParseFloat(data.KWH,10) // base 10 system 64 bit
	if err!= nil {
		panic(err.Error())
	}
	var billeddata BilledData
	if kwdata > 0 {
		billeddata.First = 40
	}
	if kwdata > 1 {
		var i float64
		billeddata.Restall = 0
		for i=1;i< kwdata;i++  {
			billeddata.Restall += 50
		}
	}


}
