package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	)

const statusSuccess  = 1
const statusFailed = 2

func FuncHandler()  {
	router := mux.NewRouter()

	router.HandleFunc("/data", GetMeterData).Methods("GET")
	router.HandleFunc("/meterdata", PostMeterData).Methods("POST")

	router.HandleFunc("/newuser", Createuser).Methods("POST")
	router.HandleFunc("/login", Getuser).Methods("PUT")

	log.Printf("Serving on :8080, Go to localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	FuncHandler()
}





