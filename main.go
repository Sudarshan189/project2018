package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)




func FuncHandler()  {
	router := mux.NewRouter()

	//router.HandleFunc("/me/{name}", MyData).Methods("GET")
	router.HandleFunc("/data", GetMeterData).Methods("GET")
	router.HandleFunc("/meterdata", PostMeterData).Methods("POST")

	// Login
	router.HandleFunc("/createbill", CreateBill).Methods("PUT")
	router.HandleFunc("/newuser", Createuser).Methods("POST")
	router.HandleFunc("/login", Getuser).Methods("PUT")
	router.HandleFunc("/user/{id}", Senduser).Methods("GET")
	router.HandleFunc("/changepass/{id}/{email}", Changepassword).Methods("POST")

	log.Printf("Serving on :8080, Go to localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}






func main() {
	go ArduinoServer()
	FuncHandler()
}





