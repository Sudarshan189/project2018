package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"log"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var (
	api_token = "sYX6pPAYT9OUg92fiNldYlhGAOfU1JwENYMRiKlNDmwb84IuLrXYqJqMR58C"
	senderid  = "BESCOM"
	number    string
	route     = "4"
)

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

	// number = string(data.Phone)
	number= strconv.FormatUint(data.Phone, 10)
	// mydata := strconv.FormatFloat(newdata.KWH*5, 'f', 10, 64)
	// fmt.Println(number)
	// message := "Your Bill Amount for RR Number"
	// resp,err := http.Get("https://www.pay2all.in/web-api/send_sms?api_token=" + api_token + "&senderid=" + senderid + "&number=" + number + "&message=" + message + "&route=" + route+"/")

	resp,err := http.Get("https://www.pay2all.in/web-api/send_sms?api_token=sYX6pPAYT9OUg92fiNldYlhGAOfU1JwENYMRiKlNDmwb84IuLrXYqJqMR58C&senderid=BESCOM&number=7795325592&message=Deepika loves sudarshan&route=4")
	if err != nil {
		panic(err)
	}
	log.Println(resp)
	//Showlog(w, r)
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//_, err1 := os.Stdout.Write(body)
	//if err1 != nil {
	//	panic(err1)
	//}
	//log.Println(resp)

}
