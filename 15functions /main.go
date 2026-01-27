package main
import (
	"fmt"
)
func main(){
	fmt.Println("welcome back lets read functions in go lang!")
	greets()
	greets2()

	fmt.Println("Result is: ", add(3,5))

	fmt.Println("Pro adder result is: ", proAdder(2,3,4,5,6,7,8,9))
}

func greets(){
	fmt.Println("hello from greets function")
}

func greets2(){
	fmt.Println("hello from greets2 function")
}

func add(a int, b int) int {
	return a + b
}


func proAdder(values...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}