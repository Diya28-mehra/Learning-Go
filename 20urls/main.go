package main

import "fmt"
import "net/url"
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"
func main(){
	fmt.Println("Welcome to handling URLs in Golang")
	fmt.Println(myurl)

	//parsing the url 
	result,_ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println("The type of query params are: %T\n",qparams)

	fmt.Println(qparams["coursename"])

	for _,val := range qparams {
		fmt.Println("Params is: ",val);
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=diya",
	}

	anotherURL:=partsOfUrl.String()
	fmt.Println(anotherURL)
}

