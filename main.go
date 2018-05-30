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

	// Create Bill for Each User
	router.HandleFunc("/createbill/{rr_num}", CreateBill).Methods("PUT")

	// Create Bill for All User
	router.HandleFunc("/createallbill", CreateAllBill).Methods("GET")

	// Get Bills with rr_num
	router.HandleFunc("/billinfo/{rr_num}", CreateBillInfo).Methods("GET")

	// Login API
	router.HandleFunc("/newuser", Createuser).Methods("POST")
	router.HandleFunc("/login", Getuser).Methods("PUT")
	router.HandleFunc("/user/{id}", Senduser).Methods("GET")
	router.HandleFunc("/changepass/{id}/{email}", Changepassword).Methods("POST")

	// Sending Email to each user
	router.HandleFunc("/sendemail/{id}", EmailService).Methods("GET")

	// Sending Email to All user with Paid == false
	router.HandleFunc("/sendemailtoall", MailtoAll).Methods("GET")

	// Sending Message to each user
	router.HandleFunc("/sendmessage/{id}", SendMessage).Methods("GET")

	// Sending Message to All user with paid == false
	router.HandleFunc("/sendmessagetoall", SendMessagetoAll).Methods("GET")

	// Get Bills with rr_num
	router.HandleFunc("/getbills/{rr_num}", GetBills).Methods("GET")

	// Paid or unpaid
	router.HandleFunc("/paid/{rr_num}/{bill_num}", PaidBill).Methods("GET")
	router.HandleFunc("/unpaid/{rr_num}/{bill_num}", UnpaidBill).Methods("GET")

	// Smart Socket
	router.HandleFunc("/smartsoc/{rr_num}/{soc_num}/{sta}", SharpThing).Methods("GET")
	router.HandleFunc("/getstate/{rr_num}", GetSwitchState).Methods("GET")


	log.Printf("Serving on :8080, Go to localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Request-Headers", "Access-Control-Request-Method"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}

func main() {
	 go ArduinoServer()
	 go SharpSoc()
	 FuncHandler()

}
