package main

import (
	"net/http"
	"strconv"
	"log"
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
)
var (
	api_token = "sYX6pPAYT9OUg92fiNldYlhGAOfU1JwENYMRiKlNDmwb84IuLrXYqJqMR58C"
	senderid  = "BESCOM"
	number    string
	route     = "4"
)


/*
For space in message add %20   .....
 */

var newbill KWHUpdate

func SendMessage(w http.ResponseWriter, r *http.Request) {
	data := Users{}
	db, err := gorm.Open("mysql", "session:session@/project2018")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	vars := mux.Vars(r)
	iddata := vars["id"]
	id, err := strconv.ParseUint(iddata,10, 64) // base 10 system 64 bit
	if err != nil {
		panic(err.Error())
	}
	db.Where(&Users{ID: uint64(id)}).First(&data)
	db.Where(&KWHUpdate{RRNum: data.RRNum}).First(&newbill)



	// number = string(data.Phone)
	number= strconv.FormatUint(data.Phone, 10)
	 mydata := strconv.FormatFloat(newdata.KWH*5, 'f', 10, 64)
	// fmt.Println(number)
	 message := "Your%20Bill%20Amount%20for%20RR%20Number%20"+data.RRNum+"%20is%20Rupees%20"+mydata+".%20Please%20pay%20before%20last%20date%20to%20avoid%20panalty.%20%20Thank%20You"


	resp,err := http.Get("https://www.pay2all.in/web-api/send_sms?api_token=sYX6pPAYT9OUg92fiNldYlhGAOfU1JwENYMRiKlNDmwb84IuLrXYqJqMR58C&senderid=BESCOM&number="+number+"&message="+message+"&route=4")
	if err != nil {
		panic(err)
	}
	log.Println(resp)
	//Showlog(w, r)
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MessageSent)


}
