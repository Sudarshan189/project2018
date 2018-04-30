package main

import (
	"net/http"


	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"strconv"

	"encoding/json"
	"log"
	"io/ioutil"
	"os"
)


var (
	api_token = "sYX6pPAYT9OUg92fiNldYlhGAOfU1JwENYMRiKlNDmwb84IuLrXYqJqMR58C"
	senderid = "BESCOM"
	number string
	message = "First+Message+from+Sudarshan"
	route = "4"


)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	data := Users{}
	db,err:= gorm.Open("mysql", "akshay:deepika019@/project2018")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	vars := mux.Vars(r)
	iddata := vars["id"]
	id,err:= strconv.ParseUint(iddata, 10,32) // base 10 system 64 bit
	if err!= nil {
		panic(err.Error())
	}
	db.Where(&Users{ID: uint(id)}).First(&data)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(data)


	number = string(data.Phone)


	resp, err:= http.Get("https://www.pay2all.in/web-api/send_sms?api_token="+ api_token + "&senderid="+senderid+"&number="+ number +"&message="+message+"&route="+route)
	if err!= nil {
		panic(err)
	}
	Showlog(w,r)
	defer resp.Body.Close()
	body, err:= ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	_,err1:= os.Stdout.Write(body)
	if err1 != nil {
		panic(err1)
	}
	log.Println(resp)


}
