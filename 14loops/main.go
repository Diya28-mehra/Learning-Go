package main 

import (
	"fmt"
)
func main(){
	fmt.Println("Welcome to loops in Go lang")

	days := []string {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("Days are:", days)

	// for d:=0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i:=range(days) {
	// 	fmt.Println(days[i])
	// }

	for index,day := range days{
		fmt.Printf("Index is %v and day is %v\n", index, day)
	}

	roguevalue:=1
	for roguevalue<10{
		fmt.Println("Rogue value is:", roguevalue)
		roguevalue+=2
	}

}
