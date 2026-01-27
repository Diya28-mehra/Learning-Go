package main 

import (
	"fmt"
)
func main(){
	fmt.Println("Welcome to pointers class ")
	// var ptr *int 
	// fmt.Println("value of pointer is ", ptr) //default value is nil 


	//for reference we use & operator
	mynum:=25
	var ptr = &mynum
	fmt.Println("value of pointer is ", ptr) 
	fmt.Println("value of pointer is ", *ptr)

	*ptr *=2
	fmt.Println("value of pointer after changing is ", *ptr)
	fmt.Println("value of mynum after changing pointer is ", mynum)
}
