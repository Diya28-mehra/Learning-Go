package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main (){
	fmt.Println("Welcome to our pizza hut")
	fmt.Println("PLease rate our pizza from 1 to 5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Printf("You rated our pizza: %s", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if (err!= nil){
		fmt.Println(err)
	} else {
		fmt.Println("You rated our pizza:", numRating+1)
	}
}