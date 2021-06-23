package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/params"
	"github.com/gorilla/mux"
)


func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Courses API")
}

func allCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of all courses")
}

func course(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprintf(w, "This is " + params["courseid"])
}


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1", home)
	router.HandleFunc("/api/v1/courses", allCourses)
	router.HandleFunc("/api/v1/courses/{courseid}", course)

	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}