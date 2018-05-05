package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/jinzhu/gorm"
)

type Mymeter struct {
	RRNum string `json:"rr_num"`
	KWH   string `json:"kwh"`
}

type KWHUpdate struct {
	RRNum     string     `gorm:"primary_key" json:"rr_num"` // Revenue Registration Number ( Unique ID )
	Current   string     `gorm:"varchar(10)" json:"current"`
	Voltage   string     `gorm:"varchar(10)" json:"voltage"`
	KWH       string     `gorm:"varchar(20)" json:"kwh"` // Consumption ( Present - Previous )
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
		_ = json.NewDecoder(conn).Decode(&meter)
		newmeter = KWHUpdate{RRNum: meter.RRNum, KWH: meter.KWH}

		db.Where(&KWHUpdate{RRNum: newmeter.RRNum}).First(&finddata)
		if finddata.RRNum != newmeter.RRNum {
			db.Create(&newmeter)
		} else {
			db.Model(&finddata).Updates(KWHUpdate{KWH: newmeter.KWH})
		}
		fmt.Println(meter)

	}

}
