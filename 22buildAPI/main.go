package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func seeding() {
	var c = Course{"1", "Go lang", 100, &Author{"Harsh sachin said", "codefreak.co.in"}}
	courses = append(courses, c)
}

// main function
func main() {

	// add raw data
	seeding()

	// routing
	r := mux.NewRouter()
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/addcourse", addOneCourse).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))
}

// controller

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Home Page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// get id from []courses
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// err
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"result": "some error occured",
	})
}

func addOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// check if body is null
	if r.Body == nil {
		json.NewEncoder(w).Encode("Body is null")
	}

	// check if received json is empty
	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Course is empty")
	}

	// add course
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}
