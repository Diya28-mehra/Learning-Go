package main

import(
	"fmt"
	"sort"
)

func main() {

	var fruitlist = []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Printf("Type of fruitlist is %T\n", fruitlist)

	fruitlist = append(fruitlist,"fig","grape")
	fmt.Println("Fuitlist ",fruitlist)

	fruitlist = append(fruitlist[1:],"kiwi")
	fmt.Println("Fuitlist ",fruitlist)

	highscore := make([]int,4)
	highscore[0] = 234
	highscore[1] = 456
	highscore[2] = 678
	highscore[3] = 890
	//highscore[4] = 1000 // This will cause an error: index out of range
	highscore = append(highscore,1000,1200,1400)
	//with append we can add more elements beyond the initial length as it allocates more memory
	sort.Ints(highscore)
	fmt.Println("Highscores ",highscore)

	//how to remove a value from slice based on index
	var courses = []string{"reactjs","javascript","swift","python","ruby"}
	fmt.Println("Courses before removing ",courses)
	var index int = 2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println("Courses after removing ",courses)

	s := "Hello, 世界"
	// Slicing the string
	substr1 := s[0:5]      // "Hello"
	substr2 := s[7:10]     // "世"
	substr3 := s[10:13]    // "界"

	fmt.Println(substr1)
	fmt.Println(substr2)
	fmt.Println(substr3)

	// Iterating over the string to show rune boundaries
	for i, r := range s {
		fmt.Printf("Index: %d, Rune: %c\n", i, r)
	}
}
