package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"net/url"
)
func main(){
	fmt.Println("welcome to web verbs")
	PerformPostFormRequest()
}


func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/postform"
	data := url.Values{}
	data.Add("firstname","diya")
	data.Add("lastname","mehra")
	data.Add("email","itsdiyamehra@gmail.com")
	data.Add("age","21")
	response,err := http.PostForm(myurl,data)
	if err!=nil{
		panic(err)
	} 
	defer response.Body.Close()
	content,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"
	//fake json payload
	requestBody := strings.NewReader(`
	{	
		"coursename":"Golearn",
		"price":100,
		"platform":"Youtube"
	}
	`)

	response,err := http.Post(myurl,"application/json",requestBody,)
	if err!=nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Status:", response.Status)

	content,_:= ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"
	response,err := http.Get(myurl)
	if err!=nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status Code: ",response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content,_ :=ioutil.ReadAll(response.Body)
	//content is in byte format
	//we need somebody who can translate
	//fmt.Println("Content is: ",string(content))

	var responseString strings.Builder 

	bytecount,_:=responseString.Write(content)
	fmt.Println("ByteCount is: ",bytecount)
	fmt.Println("Content is: ",responseString.String())
}


