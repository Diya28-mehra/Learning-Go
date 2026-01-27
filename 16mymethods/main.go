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
	diya.Getstatus()
	diya.NewMail()
	fmt.Println("Original mail: ",diya.Email)
}

type User struct {
	Name string 
	Email string 
	Status bool
	Age int 
}


func (u User) Getstatus(){
	fmt.Println("Is User active? ",u.Status)
}


func (u User) NewMail(){
	u.Email = "testgo@dev"
	fmt.Println("Email of this user is: ",u.Email)
}