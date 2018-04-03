package main

import (
	"net"
	"fmt"
	"encoding/json"
	//"time"
	// "github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm"
	"time"
)

type Mymeter struct {
	RRNum string `json:"rr_num"`
	KWH	string `json:"kwh"`
}

type KWHUpdate struct {

	RRNum 		string 		`gorm:"primary_key" json:"rr_num"`					// Revenue Registration Number ( Unique ID )
	KWH 		string 		`gorm:"varchar(20)" json:"kwh"`			// Consumption ( Present - Previous )
	UpdatedAt *time.Time		`json:"updated_at"`

}


func ArduinoServer() {
	var meter Mymeter
	var newmeter KWHUpdate
	var finddata KWHUpdate
	ln,_:=net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()
	 db, err := gorm.Open("mysql", "akshay:deepika019@/project2018")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&KWHUpdate{})

	for {
		_ = json.NewDecoder(conn).Decode(&meter)
		newmeter = KWHUpdate{RRNum:meter.RRNum, KWH:meter.KWH}

		db.Where(&KWHUpdate{RRNum:newmeter.RRNum}).First(&finddata)
		if finddata.RRNum != newmeter.RRNum {
			db.Create(&newmeter)
		} else {
			db.Model(&finddata).Updates(KWHUpdate{KWH:newmeter.KWH})
		}
		fmt.Println(meter)
	}

}