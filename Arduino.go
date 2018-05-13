package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type Mymeter struct {
	RRNum   string `json:"rr_num"`
	Current string `json:"current"`
	Voltage string `json:"voltage"`
	KWH     string `json:"kwh"`
}

type KWHUpdate struct {
	RRNum     string     `gorm:"primary_key" json:"rr_num"` // Revenue Registration Number ( Unique ID )
	Current   string     `gorm:"varchar(30)" json:"current"`
	Voltage   string     `gorm:"varchar(30)" json:"voltage"`
	KWH       float64    `json:"kwh"` // Consumption ( Present - Previous )
	UpdatedAt *time.Time `json:"updated_at"`
}


func ArduinoServer() {
	var meter Mymeter
	var newmeter KWHUpdate
	var finddata KWHUpdate
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()
	db, err := gorm.Open("mysql", "session:session@/project2018")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&KWHUpdate{})

	for {
		_ = json.NewDecoder(conn).Decode(&meter) // KWH to 0 for error removal
		newmeter = KWHUpdate{RRNum: meter.RRNum, KWH: 0, Current: meter.Current, Voltage: "229.9"}
		//fmt.Println(newmeter)
		db.Where(&KWHUpdate{RRNum: newmeter.RRNum}).First(&finddata)
		if finddata.RRNum != newmeter.RRNum {
			db.Create(&newmeter)
			finddata = newmeter
		} else {



			//// Algorithm to calculate power in kWH
			cur, err := strconv.ParseFloat(finddata.Current, 10) // base 10 system 64 bit
			// Current in Amps
			if err != nil {
				panic(err.Error())
			}

			vol, err := strconv.ParseFloat(finddata.Voltage, 10) // base 10 system 64 bit
			// Voltage in Volts
			if err != nil {
				panic(err.Error())
			}

			if cur >= 3 {
				cur = 0
			}
			power := (cur * vol) / 1000 // Power is Voltage times Current in kW


			/*
			Now we need to calculate kw in kwh..
			This need small calculation.
			if server is updating per 0.5 Sec then divide it with 3600 to conver it in kWh
		*/
		    // divide = 2/3600
			kWh := (power*2)/3600
			//
			// //finddata.KWH = strconv.FormatFloat(kWh, 'f', 2, 64)
			//
			kwdata := finddata.KWH

			//
			finddata.KWH = kwdata + kWh


			//finddata.KWH = strconv.FormatFloat(kwdata + kWh, 'f', 2, 64)
			// add kWh data to previous value
			//
			// fmt.Println(finddata.KWH)
			//
			//
			db.Model(&finddata).Updates(KWHUpdate{KWH: finddata.KWH, Current: newmeter.Current, Voltage: newmeter.Voltage})

			// db.Model(&finddata).Updates(KWHUpdate{})
			//fmt.Print(finddata.KWH)
			//
			// _ = json.NewEncoder(conn).Encode(&finddata.KWH)
			fmt.Println(finddata)
		}
	}
}
