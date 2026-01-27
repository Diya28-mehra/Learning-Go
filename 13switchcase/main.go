package main 

import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	// switch case statement
	fmt.Println("Welcome to switch case in Go")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice number is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 You can open")
	case 2:
		fmt.Println("Dice value is 2 You can move 2 steps")
	case 3:
		fmt.Println("Dice value is 3 You can move 3 steps")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4 You can move 4 steps")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5 You can move 5 steps")
	case 6:
		fmt.Println("Dice value is 6 You can move 6 steps and roll the dice again")
	default:
		fmt.Println("Invalid Dice number")
	}
}
