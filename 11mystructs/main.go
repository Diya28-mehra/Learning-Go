package main 

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome to structs in Go lang")

	//no inheritance in Go lang: No super or parent 

	diya := User{"Diya","itsdiyamehra@gmail.com",true,21}
	fmt.Println(diya)
	fmt.Printf("Diya details are: %+v\n",diya)
	fmt.Printf("Name is %v and email is %v\n",diya.Name,diya.Email)
}

type User struct {
	Name string 
	Email string 
	Status bool
	Age int 
}

