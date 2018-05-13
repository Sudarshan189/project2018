package main

import (
	"encoding/json"
	"net/http"
	"net/smtp"
	"strconv"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var (
	newdata KWHUpdate
)

func EmailService(w http.ResponseWriter, r *http.Request) {

	hostUrl := "smtp.gmail.com"
	hostPort := "587"
	emailSender := "cmrroutine@gmail.com"
	passWord := "Sodium11"
	data := Users{}
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
	emailAuth := smtp.PlainAuth("", emailSender, passWord, hostUrl)
	mydata := strconv.FormatFloat(newdata.KWH*5, 'f', 10, 64)
	msg := []byte("To:" + data.Email + "\r\n" +
		"Subject: BESCOM Electricity Bill \r\n" +
		"\r\n" +
		"RR Num: " + data.RRNum + "\r\n" + "Name: " + data.Name + "\r\n" + "Address: " + data.Address + "\r\n" + "Email: " + data.Email + "\r\n" + "\r\n" + "Bill Amount: " + mydata + "\r\n")
	fmt.Println(newdata.KWH)
	err1 := smtp.SendMail(hostUrl+":"+hostPort, emailAuth, emailSender, []string{data.Email}, msg)
	if err1 != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(EMailFailed)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(EmailSent)
	}
}

func MailtoAll(w http.ResponseWriter, r *http.Request) {
	var newbills []Bill
	hostUrl := "smtp.gmail.com"
	hostPort := "587"
	emailSender := "cmrroutine@gmail.com"
	passWord := "Sodium11"
	var data []Users
	db, err := gorm.Open("mysql", "session:session@/project2018")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	Showlog(w, r)
	link := "http://sarkarinews.com/wp-content/uploads/2015/02/bescom.jpg"
	db.Where(&Users{}).Find(&data)
	for i := 0; i < len(data); i++ {
		db.Where(&Bill{RRNum: data[i].RRNum}).Find(&newbills)
		for i := 0; i < len(newbills); i++ {
			if newbills[i].Paid == false {
				emailAuth := smtp.PlainAuth("", emailSender, passWord, hostUrl)
				mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
				subject := "Subject: BESCOM Electricity Bill\n"
				msg := []byte(subject + mime + "<html><body style=\"border: 2px solid black; margin:8px;\"><div align=\"center\"><img src=" + link + " style=\"width:100px;height:125px;\"></div><div align=\"center\"><h2>Name: " + newbills[i].Name + " </h2><h3>Address: " + newbills[i].Address + "</h3>" +
					"<h3> Bill Number: " + fmt.Sprint(newbills[i].BillNum) + "</h3><br><h3>Email: " + newbills[i].Email + "</h3>" +
					"<h3>Phone: " + fmt.Sprint(newbills[i].Phone) + "</h3><br><h3> Bill Amount: " + fmt.Sprint(newbills[i].NetAmtDue) + "</h3><h3> Due Date: " + newbills[i].DueDate + "</h3> <button style=\"background-color: #4CAF50;color: white;padding: 15px 32px;text-align: center;text-decoration: none;display: inline-block;font-size: 16px;margin: 4px 2px;cursor: pointer;\"><a style=\"color: white;\" href=\"http://localhost:4200/login\">Pay Bill</a></button></div></body></html>")
				err1 := smtp.SendMail(hostUrl+":"+hostPort, emailAuth, emailSender, []string{newbills[i].Email}, msg)
				if err1 != nil {
					fmt.Println("Mail failed to " + newbills[i].Email)
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(EMailFailed)
				} else {
					fmt.Println("Mail sent to " + newbills[i].Email)
				}

			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(EmailSent)
}
