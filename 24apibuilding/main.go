package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model for course - file
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}


//fake db
var courses []Course


//middleware, helper - file
func (c* Course) IsEmpty() bool{
	return  c.CourseName==""
}

func main(){
	fmt.Println("Welcome to build api")
	r:=mux.NewRouter()
	//seeding
	courses = append(courses, Course{CourseId: "23",CourseName: "ReactJS", CoursePrice: 723,
            Author: &Author{Fullname: "Diya Mehra",Website: "www.diyacourse.com"}})

	courses = append(courses, Course{CourseId: "12",CourseName: "Angular", CoursePrice: 723,
            Author: &Author{Fullname: "Kims javy",Website: "www.courses.com"}})

	courses = append(courses, Course{CourseId: "56",CourseName: "Go Lang", CoursePrice: 723,
            Author: &Author{Fullname: "John Dave",Website: "www.seedaves.com"}})

	courses = append(courses, Course{CourseId: "45",CourseName: "Blockchain", CoursePrice: 723,
            Author: &Author{Fullname: "Kia Kapoor",Website: "www.kapoorsstudies.com"}})

	//routing
	r.HandleFunc("/",ServeHome).Methods("GET")
	r.HandleFunc("/courses",GetAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}",GetOneCourse).Methods("GET")
	r.HandleFunc("/course/",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateCourse).Methods("PUT")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000",r))

}

//controllers - in different file or folder

//serve home route
func ServeHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1> Welcome to API learning in Go Lang Ha Ha Ha</h1>"))
}


func GetAllCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type","application/json")

	//grab id from request
	params:=mux.Vars(r)

	//loop through courses, find matching id and return the response
	for _,course := range courses{
		if course.CourseId==params["id"]{
			json.NewEncoder(w).Encode(course)
			return 
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("create One Course")
	w.Header().Set("Content-Type","application/json")


	//what if body is empty
	if r.Body==nil{
		json.NewEncoder(w).Encode("Please send some Data")
		return
	}

	//what about {}
	var course Course
	_=json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return 
	}

	//generate course id , convert it to string 
	//append new course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId=strconv.Itoa(rand.Intn(100))
	courses = append(courses,course)
	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Updating Course")
	w.Header().Set("Content-Type","application/json")

	//first grab the id from request 
	params:=mux.Vars(r)

	//loop through the courses, fetch the id, remove it and add again
	for index,course := range courses{
		if course.CourseId==params["id"]{
			courses = append(courses[:index],courses[index+1:]...)
			var course Course
			_=json.NewDecoder(r.Body).Decode(&course)
			course .CourseId=params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//to do send a reponse when id is not found 
	json.NewEncoder(w).Encode("No course found with given id")
}

func deleteCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete a Course")
	w.Header().Set("Content-Type","application/json")

	params:=mux.Vars(r)

	for index,course:=range courses{
		if course.CourseId==params["id"]{
			courses = append(courses[:index],courses[index+1:]...)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}