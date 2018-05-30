package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
)

/*
For space in message add %20   .....
*/

var (
	newbill KWHUpdate
	number  string
	rr string
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	data := Users{}
	db, err := gorm.Open("mysql", "session:session@/project2018")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	vars := mux.Vars(r)
	iddata := vars["id"]
	id, err := strconv.ParseUint(iddata, 10, 64) // base 10 system 64 bit
	if err != nil {
		panic(err.Error())
	}
	db.Where(&Users{ID: uint64(id)}).First(&data)
	db.Where(&KWHUpdate{RRNum: data.RRNum}).First(&newbill)
	number = strconv.FormatUint(data.Phone, 10)
	mydata := strconv.FormatFloat(newdata.KWH*5, 'f', 10, 64)
	message := "Your%20Bill%20Amount%20for%20RR%20Number%20" + data.RRNum + "%20is%20Rupees%20" + mydata + ".%20Please%20pay%20before%20last%20date%20to%20avoid%20panalty.%20%20Thank%20You"
	resp, err := http.Get("https://www.pay2all.in/web-api/send_sms?api_token=sYX6pPAYT9OUg92fiNldYlhGAOfU1JwENYMRiKlNDmwb84IuLrXYqJqMR58C&senderid=BESCOM&number=" + number + "&message=" + message + "&route=4")
	if err != nil {
		panic(err)
	}
	log.Println(resp)
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MessageSent)
}

func SendMessagetoAll(w http.ResponseWriter, r *http.Request) {
	var data []Users
	var newbills []Bill
	db, err := gorm.Open("mysql", "session:session@/project2018")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Where(&Users{}).Find(&data)
	for i := 0; i < len(data); i++ {
		db.Where(&Bill{RRNum: data[i].RRNum}).Find(&newbills)
		for i := 0; i < len(newbills); i++ {
			if newbills[i].Paid == false {
				message := "Your%20Bill%20Amount%20for%20RR%20Number%20" + newbills[i].RRNum + "%20is%20Rupees%20" + fmt.Sprint(newbills[i].NetAmtDue) + ".%20Please%20pay%20before%20" + newbills[i].DueDate + "%20to%20avoid%20penalty.%20%20Thank%20You"
				resp, err := http.Get("https://www.pay2all.in/web-api/send_sms?api_token=sYX6pPAYT9OUg92fiNldYlhGAOfU1JwENYMRiKlNDmwb84IuLrXYqJqMR58C&senderid=BESCOM&number=" + fmt.Sprint(newbills[i].Phone) + "&message=" + message + "&route=4")
				if err != nil {
					panic(err)
				}
				log.Println(resp)
				//Showlog(w, r)
				defer resp.Body.Close()
				fmt.Println("Message Sent to " + newbills[i].Email)
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MessageSent)
}


func SendLoadMsg(rr string) {
	data := Users{}
	db, err := gorm.Open("mysql", "session:session@/project2018")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//vars := mux.Vars(r)
	//iddata := vars["id"]
	//id, err := strconv.ParseUint(iddata, 10, 64) // base 10 system 64 bit
	//if err != nil {
	//	panic(err.Error())
	//}
	db.Where(&Users{RRNum: rr}).First(&data)
	// db.Where(&KWHUpdate{RRNum: data.RRNum}).First(&newbill)
	 number = strconv.FormatUint(data.Phone, 10)
	// mydata := strconv.FormatFloat(newdata.KWH*5, 'f', 10, 64)
	message := "Load%20Exceeded.%20Supply%20has%20been%20stopped%20for%20rrnum%20"+rr+"%20";
	resp, err := http.Get("https://www.pay2all.in/web-api/send_sms?api_token=sYX6pPAYT9OUg92fiNldYlhGAOfU1JwENYMRiKlNDmwb84IuLrXYqJqMR58C&senderid=BESCOM&number=" + number + "&message=" + message + "&route=4")
	if err != nil {
		panic(err)
	}
	log.Println(resp)
	defer resp.Body.Close()
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(MessageSent)
}

