package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("------------------------")

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	// The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value.
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	fruits2 := [3]string{"apple", "orange", "banana"}
	for _, val := range fruits2 {
		fmt.Printf("%v\n", val)
	}
}
