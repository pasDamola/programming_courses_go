package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type courseInfo struct {
	Title string `json:"Title"`
}

var courses map[string]courseInfo

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Courses API")
}

func allCourses(w http.ResponseWriter, r *http.Request) {
	kv := r.URL.Query()
	for i, v := range kv {
		fmt.Println(i, v)
	}
	fmt.Fprintf(w, "List of all courses ")

	// returns all the courses in JSON
	json.NewEncoder(w).Encode(courses)

}

func course(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// fmt.Fprintf(w, "This is "+params["courseid"])
	// fmt.Fprintf(w, "\n")
	// fmt.Fprintf(w, r.Method)
	if r.Header.Get("Content-type") == "application/json" {
		// POST is for creating new course
		if r.Method == "POST" {
			// read the string sent to the service
			var newCourse courseInfo
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil {
				// convert JSON to object
				json.Unmarshal(reqBody, &newCourse)
				if newCourse.Title == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply course " + "information " + "in JSON format"))
					return
				}
				// check if course exists; add only if
				// course does not exist
				if _, ok := courses[params["courseid"]]; !ok {
					courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " + params["courseid"]))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte("409 - Duplicate course ID"))
				}
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply course information " + "in JSON format"))
			}
		}

		//---PUT is for creating or updating
		// existing course---
		if r.Method == "PUT" {
			var newCourse courseInfo
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil {
				json.Unmarshal(reqBody, &newCourse)
				if newCourse.Title == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply course " + " information " + "in JSON format"))
					return
				}
				// check if course exists; add only if
				// course does not exist
				if _, ok := courses[params["courseid"]]; !ok {
					courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " + params["courseid"]))
				} else {
					// update course
					courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusNoContent)
				}
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply " + "course information " + "in JSON format"))
			}
		}
	}
}

func main() {
	router := mux.NewRouter()

	courses = make(map[string]courseInfo)
	router.HandleFunc("/api/v1", home)
	router.HandleFunc("/api/v1/courses", allCourses)
	router.HandleFunc("/api/v1/courses/{courseid}", course).Methods("GET", "POST", "PUT", "DELETE")
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
