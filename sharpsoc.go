package main

import (
	"fmt"
	"net"
	"encoding/json"
	// "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)



type Status struct {
	State    string `json:"state"`
}



//
//type Smart struct {
//	RRNum string `json:"rr_num"`
//	SocketId  string `json:"socket_id"`
//	SocStatus bool   `json:"soc_status"`
//}
//
//var (
//	socket Socket
//
//)
//func Sharpsoc(w http.ResponseWriter, r *http.Request) {
//	db,err := gorm.Open("mysql", "session:session@/project2018")
//	if err!= nil {
//		panic(err)
//	}
//	defer db.Close()
//	db.AutoMigrate(&Socket{})
//	Showlog(w, r)
//	_=json.NewDecoder(r.Body).Decode(&socket)
//	fmt.Println(socket)
//
//	db.Where(&Socket{RRNum:socket.RRNum, SocketId:socket.SocketId}).First()
//
//
//}



func SharpSoc() {
	var sta Status
	var found []Socket
	db,err := gorm.Open("mysql", "session:session@/project2018")
		if err!= nil {
			panic(err)
		}
	defer db.Close()
	// db.AutoMigrate(&Socket{})

	ln, _ := net.Listen("tcp", ":8082")
	conn, _ := ln.Accept()

	for  {
		db.Where(&Socket{}).Find(&found)
		//fmt.Println(found)
		if found[0].Limit == "0" {
			one:= found[0].RRNum+":"+found[0].SocketId+":"+found[0].SocStatus
			fmt.Fprint(conn,one)
			_ = json.NewDecoder(conn).Decode(&sta) // KWH to 0 for error removal
			fmt.Println(sta)
		} else {
			db.Model(&found[0]).Update("soc_status", "0")
			one:= found[0].RRNum+":"+found[0].SocketId+":"+found[0].SocStatus
			fmt.Fprint(conn,one)
			_ = json.NewDecoder(conn).Decode(&sta) // KWH to 0 for error removal
			fmt.Println(sta)
		}

		if found[1].Limit == "0" {
			two := found[1].RRNum + ":" + found[1].SocketId + ":" + found[1].SocStatus
			fmt.Fprint(conn, two)
			_ = json.NewDecoder(conn).Decode(&sta) // KWH to 0 for error removal
			fmt.Println(sta)
		} else {
			db.Model(&found[1]).Update("soc_status", "0")
			one:= found[1].RRNum+":"+found[1].SocketId+":"+found[1].SocStatus
			fmt.Fprint(conn,one)
			_ = json.NewDecoder(conn).Decode(&sta) // KWH to 0 for error removal
			fmt.Println(sta)
		}
	}
}