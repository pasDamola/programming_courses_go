package main

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
)


func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Courses API")
}


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", home)

	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}