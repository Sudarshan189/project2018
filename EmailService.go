package main

import (

	"net/http"
	"net/smtp"
	"strconv"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"fmt"
)



var newdata KWHUpdate

func EmailService(w http.ResponseWriter, r *http.Request) {

	hostUrl := "smtp.gmail.com"
	hostPort := "587"
	emailSender := "cmrroutine@gmail.com"
	passWord := "Sodium11"
	// emailReceiver := "shanbhagsudharshan@gmail.com"
	data := Users{}
	// Data from databases
	db, err := gorm.Open("mysql", "session:session@/project2018")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	Showlog(w, r)
	vars := mux.Vars(r)
	id := vars["id"]
	iddata, err := strconv.ParseUint(id, 10, 64) // base 10 system 64 bit
	if err != nil {
		panic(err.Error())
	}

	db.Where(&Users{ID: uint64(iddata)}).First(&data)

	db.Where(&KWHUpdate{RRNum: data.RRNum}).First(&newdata)
	// fmt.Println(newdata)

	emailAuth := smtp.PlainAuth("", emailSender, passWord, hostUrl)

	// mydata = strconv.FormatFloat(newdata.KWH*5, 'f', 2, 64)
	mydata := strconv.FormatFloat(newdata.KWH*5, 'f', 10, 64)
	// phonedata := strconv.FormatUint(data.Phone, 64)

	msg := []byte("To:" + data.Email + "\r\n" +
		"Subject: BESCOM Electricity Bill \r\n" +
		"\r\n" +
		 "RR Num: "+data.RRNum + "\r\n" + "Name: "+data.Name+ "\r\n"+ "Address: "+data.Address + "\r\n" +"Email: "+data.Email + "\r\n"+ "\r\n"+"Bill Amount: "+mydata + "\r\n")

	 fmt.Println(newdata.KWH)



	// fmt.Println(mydata)

	 err1 := smtp.SendMail(hostUrl+":"+hostPort, emailAuth, emailSender, []string{data.Email}, msg)

	if err1 != nil {
		//fmt.Print("Error")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(EMailFailed)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(EmailSent)
	}


}
