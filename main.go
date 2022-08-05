package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	}).Methods("GET")
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	}).Methods("GET")
	router.HandleFunc("/re1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "re1")
	}).Methods("GET")
	http.ListenAndServe(":8080", router)
}
