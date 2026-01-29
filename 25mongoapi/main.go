package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Diya28-mehra/router"
)

func main(){
	fmt.Println("Mongo DB API")
	fmt.Println("Server is getting started...")
	r:=router.Router()
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Listening at port 4000...")
}
