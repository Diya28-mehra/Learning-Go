package main

import (
	"fmt"
)
func main (){
	fmt.Println("Hello")
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("oye hoye")
	myDefer()
}

//order will be first in last out 
//hello
//ooye hoye
//4 3 2 1 0
//two one world

func myDefer(){
	for i:=0; i<5; i++{
		defer fmt.Println(i);
	}
}