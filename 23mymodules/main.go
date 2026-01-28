package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)
func main(){
	fmt.Println("Hello mod in Go lang")
	greeter()
	r:=mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000",r))
}

func greeter(){
	fmt.Println("Hey there are mod users")
}

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome to Go Lang series on Youtube</h1>"))
}