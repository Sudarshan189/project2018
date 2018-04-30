package main

import (
	"net/http"
	"net/smtp"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
)

func EmailService(w http.ResponseWriter, r *http.Request) {

	hostUrl :="smtp.gmail.com"
	hostPort := "587"
	emailSender:= "cmrroutine@gmail.com"
	passWord := "Sodium11"
	 // emailReceiver := "shanbhagsudharshan@gmail.com"
	data := Users{}
	// Data from databases
	db, err := gorm.Open("mysql", "akshay:deepika019@/project2018")
	if err!= nil {
		panic(err.Error())
	}
	defer db.Close()
	Showlog(w,r)
	vars := mux.Vars(r)
	id := vars["id"]
	iddata,err:= strconv.ParseUint(id, 10,32) // base 10 system 64 bit
	if err!= nil {
		panic(err.Error())
	}

	db.Where(&Users{ID:uint(iddata)}).First(&data)




	emailAuth := smtp.PlainAuth("", emailSender, passWord, hostUrl)

	msg := []byte("To:"+data.Email+"\r\n" +
		"Subject: BESCOM Electricity Bill \r\n" +
		"\r\n" +
		"saj"+"\r\n")

	err1 := smtp.SendMail(hostUrl+ ":"+ hostPort, emailAuth, emailSender, []string{data.Email}, msg)

	if err1 != nil {
		fmt.Print("Error")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(EmailSent)


}

