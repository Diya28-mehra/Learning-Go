package main 

import "fmt"


const LoginTOken = "kfsjdfjafasjf" //captial letter in starting makes the variable public 

func main(){
	fmt.Println("Variables in Go")
	var name string = "Diya"
	fmt.Println(name)
	fmt.Printf("variable :%T\n", name)

	var isLoggedIn bool = false;
	fmt.Println(isLoggedIn)
	fmt.Printf("variable :%T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable :%T\n", smallVal)

	//default values and some aliases
	var num int
	fmt.Println(num)
	fmt.Printf("variable :%T\n", num)

	var s string 
	fmt.Println(s)
	fmt.Printf("variable :%T\n", s)

	var floatNum float32 = 255.455
	fmt.Println(floatNum)
	fmt.Printf("variable :%T\n", floatNum)

	//implicit typeing
	//that is where lexer and parser decides the type by the value assigned
	var city = "Bangalore"
	fmt.Println(city)
	fmt.Printf("variable :%T\n", city)

	//no var style
	country := "India"
	fmt.Println(country)
	fmt.Printf("variable :%T\n", country)
}