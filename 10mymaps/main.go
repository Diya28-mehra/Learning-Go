package main

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome to maps in Go lang")
	languages := make(map[string]string) 
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["GO"] = "Golang"

	fmt.Println("List of languages: ", languages)
	fmt.Println("Length of languages map: ", len(languages))
	delete(languages,"PY")
	fmt.Println("List of languages after deletion: ", languages)

	//loops are interesting in maps 
	for key, value := range languages{
		fmt.Printf("For key %v, Value is %v \n",key,value)
	}
}
