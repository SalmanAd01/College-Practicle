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
	router.HandleFunc("/merge", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "confilct")
	}).Methods("GET")
	http.ListenAndServe(":8080", router)
}
