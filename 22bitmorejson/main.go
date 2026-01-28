package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"`
	Price int 
	Platform string `json:"website"`
	Password string `json:"-"` //entirely remove password 
	Tags []string `json:"tags,omitempty"` //omitempty
}

func EncodeJson(){
	lcoCourses := []course {
		{"ReactJS BootCamp",299,"Youtube.com","abc123",[]string {"Learn with practical examples","webdev"}},
		{"Namaste Dev",600,"www.code.com","abc234",[]string{"full stack","mern specialised"}},
		{"Namaste Angular",299,"www.code.com","ab2344",[]string{"Angular learn","easy for you"}},
		{"AI learn",600,"www.code.com","ab2344",nil},
	}

	//package this data as json data

	finalJson,err := json.MarshalIndent(lcoCourses,"","\t")
	if err!=nil{
		panic(err)
	}

	fmt.Printf("Final Json is %s",finalJson)

}
func main(){
	fmt.Println("Welcome to learn about json in Go")
	//EncodeJson()
	DecodeJson()
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
		[
			{
				"coursename": "ReactJS BootCamp",
				"price": 299,
				"website": "Youtube.com",
				"tags": [
					"Learn with practical examples",
					"webdev"
				]
			},
			{
				"coursename": "Namaste Dev",
				"price": 600,
				"website": "www.code.com",
				"tags": [
					"full stack",
					"mern specialised"
				]
			},
			{
				"coursename": "Namaste Angular",
				"price": 299,
				"website": "www.code.com",
				"tags": [
					"Angular learn",
					"easy for you"
				]
			}
		]
	`)
	var lcoCourses []course 
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid{
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb,&lcoCourses)
		fmt.Printf(" ha ha ha %#v\n",lcoCourses)
	}else{
		fmt.Println("Json was not valid")	
	}

	//some cases where you just to add data to key value 
	var myOnlineData []map[string]interface{}
	json.Unmarshal(jsonDataFromWeb,&myOnlineData)
	fmt.Printf("%#v\n",lcoCourses)

	for k,v := range myOnlineData{
		fmt.Printf("Key is %v and Value is %v and Type is: %T\n",k,v,v)
	}
}

