package main 
import (
	"fmt"

)

func main(){
	fmt.Println("Welcome to if else in Go lang")
	loggedIn := 10
	var result string 
	if loggedIn<10{
		result = "Regular User"
	}else if loggedIn>10 {
		result = "Watch out"
	}else{
		result = "Admin User"
	}
	fmt.Println(result)

	if num:=3; num<5{
		fmt.Println("Number is less than 5")
	}else{
		fmt.Println("Number is greater than 5")
	}
}
