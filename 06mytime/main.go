package main
import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Welcome to Time Study in Go lang")
	fmt.Println("Time study is important for productivity")

	presentTime := time.Now()
	fmt.Println("Present Time: ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.March, 8, 23, 0, 0, 0, time.UTC)
	fmt.Println("Created Date: ", createdDate)

}