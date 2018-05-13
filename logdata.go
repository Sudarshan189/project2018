package main

import (
	"log"
	"net/http"
)

var (
	proto      = ""
	requestURI = ""
	method     = ""
)

func Showlog(w http.ResponseWriter, r *http.Request) {
	proto, requestURI, method = r.Proto, r.RequestURI, r.Method
	log.Println(proto, requestURI, method)
}
