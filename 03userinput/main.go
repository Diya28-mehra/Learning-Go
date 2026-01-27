package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	welcomeMessage := "Welcome to Go Programming!"
	fmt.Print(welcomeMessage)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter your rating for our class: ")

	//comma ok syntax to handle error
	input , _ := reader.ReadString('\n')
	fmt.Printf("You entered: %s", input)
}