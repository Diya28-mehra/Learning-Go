package main 

import (
	"fmt"
)
func main(){
	fmt.Println("Welcome to arrays in Go lang")
	var arr[4]int 
	arr[0]=34
	arr[1]=11
	//arr[2]=56 //by default it will be 0
	arr[3]=89
	fmt.Println("Array is: ",arr)


	var veglist=[5] string {"potato","tomato","","cabbage"}
	fmt.Println("Vegetable list: ",veglist)
	fmt.Println("Length of array is: ",len(veglist))

}
