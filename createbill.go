package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var (
	slab1 float32 =3.25
	slab2 float32 =4.70
	slab3 float32 =6.25
	slab4 float32 =7.30
	slab5 float32 =7.35
	slab6 float32 =7.40

)

type BilledData struct {
	First   uint32
	Restall uint32
}

func CreateBillInfo(w http.ResponseWriter, r *http.Request) {
	var data KWHUpdate
	var bill Bill
	db, err := gorm.Open("mysql", "session:session@/project2018")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	Showlog(w, r)
	vars := mux.Vars(r)
	rrnum := vars["rr_num"]

	db.Where(&KWHUpdate{RRNum: rrnum}).First(&data)
	kwdata, err := strconv.ParseFloat(data.KWH, 10) // base 10 system 64 bit
	if err != nil {
		panic(err.Error())
	}

	var billeddata BilledData
	if kwdata > 0 {
		billeddata.First = 40
		bill.FixedCharges.FirstKW = string(40)
	}
	if kwdata > 1 {
		var i float64
		billeddata.Restall = 0
		for i = 1; i <= kwdata; i++ {
			billeddata.Restall += 50
		}
		bill.FixedCharges.EveryAddKW = string(billeddata.Restall)
	}

	/*
	DC amps to kilowatts calculation formula
	The power P in kilowatts is equal to the current I in amps, times the voltage V in volts divided by 1000:

	==>     P(kW) = I(A) × V(V) / 1000

	So kilowatts are equal to amps times volts divided by 1000:
    ==> 	kilowatt = amp × volt / 1000
	or
	==>     kW = A × V / 1000

	 */


	 /*
	 AC single phase amps to kilowatts calculation formula
	The real power P in kilowatts is equal to the power factor PF times the phase current I in amps, times the RMS voltage V in volts divided by 1000:

	==>   P(kW) = PF × I(A) × V(V) / 1000

	So kilowatts are equal to power factor times amps times volts divided by 1000:
	 ==>  kilowatt = PF × amp × volt / 1000
	 			or
	 ==>  kW = PF × A × V / 1000

	 */

	 /*
	 R.M.S. value of A.C. voltage in Indian households is 220 to 240 volts.
	 The peak value of A.C. voltage is 1.414 times the R.M.S.
	  */










}
