package main

import (
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func FuncHandler() {
	router := mux.NewRouter()

	// Meter data retrieving API
	router.HandleFunc("/details/{rr_num}", GetMeterData).Methods("GET")

	// Billing Information
	router.HandleFunc("/createbill/{rr_num}", CreateBill).Methods("PUT")
	router.HandleFunc("/createallbill", CreateAllBill).Methods("GET")
	router.HandleFunc("/billinfo/{rr_num}", CreateBillInfo).Methods("GET")

	// Login
	router.HandleFunc("/newuser", Createuser).Methods("POST")
	router.HandleFunc("/login", Getuser).Methods("PUT")
	router.HandleFunc("/user/{id}", Senduser).Methods("GET")
	router.HandleFunc("/changepass/{id}/{email}", Changepassword).Methods("POST")

	// Sending Email
	router.HandleFunc("/sendemail/{id}", EmailService).Methods("GET")

	// Sending Email to All
	router.HandleFunc("/sendemailtoall", MailtoAll).Methods("GET")

	// Sending Message
	router.HandleFunc("/sendmessage/{id}", SendMessage).Methods("GET")

	// Get Bills
	router.HandleFunc("/getbills/{rr_num}", GetBills).Methods("GET")

	// Paid or unpaid
	router.HandleFunc("/paid/{rr_num}/{bill_num}", PaidBill).Methods("GET")
	router.HandleFunc("/unpaid/{rr_num}/{bill_num}", UnpaidBill).Methods("GET")

	log.Printf("Serving on :8080, Go to localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Request-Headers", "Access-Control-Request-Method"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))


}

func main() {
	go ArduinoServer()
	FuncHandler()

}
