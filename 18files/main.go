package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	fmt.Println("Welcome to files in Go lang")

	content:= "this needs to go in file"

	file,err := os.Create("./mylocogofile.txt")

	if err!=nil{
		panic(err)
	}
	length, err := io.WriteString(file,content)
	if err!=nil{
		panic(err)
	}
	fmt.Println("Length is: ",length)
	defer file.Close()
	readFile("./mylocogofile.txt")
}


func readFile(filename string ){
	databyte, err := os.ReadFile(filename)
	if err!=nil{
		panic(err)
	}
	fmt.Println("Text data inside the file is \n", string(databyte))
}
