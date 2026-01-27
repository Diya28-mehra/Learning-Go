package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.incrementors.com/tools/bulk-url-opener/"
func main(){
	fmt.Println("LCO web requests")
	
	response,err := http.Get(url)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n",response)
	defer response.Body.Close() //callers responsibility to close connection 

	dataBytes,err := ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	content:= string(dataBytes)
	fmt.Println("content is ", content)

}
