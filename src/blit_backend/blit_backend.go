/*
 * File:    blit_backend.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	1-7-21
 *
 * Summary of File:
 *
 *  This program runs a backend server for Blit package to handle http requests from a react frontend
 *
 */


package blit_backend

import (
	"net/http"
	"github.com/gorilla/mux"
)

type FrontReq struct {
	PathReq string
	URL string
}

// Starts a backend that starts listening to http://localhost:8080
func Start() {
	router := mux.NewRouter()

	front := &FrontReq{}

	router.HandleFunc("/", front.handler)

	http.Handle("/", router)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func (front *FrontReq) handler(writter http.ResponseWriter, request *http.Request) {
	writter.WriteHeader(200)
	writter.Write([]byte("Welcome to Blit: Let's list some folders!"))
}