package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"tergo/api"
)

func importENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	importENV()
	r := mux.NewRouter()
	r.HandleFunc("/", api.FCMHandler).Methods("POST")
	http.ListenAndServe(":80", r)
}
