package main

import (
	"fmt"
	"math/big"
	"crypto/rand"
)

func main(){
	fmt.Println("Welcome to Maths in Go lang")
	var mynum1 int = 2;
	var mynum2 float64 = 3.99

	fmt.Println("The sum of these two numbers ", mynum1 + int(mynum2))


	//random number using cryptography algorithm

	//USING MATH LIBRARY
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))


	//using crypto/rand library

	myrand,_ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myrand)
}

