package main

import (
	"fmt"
)

func main() {
	// with defined length
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	// with inferred length
	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)
	fmt.Println(cars[0])
	fmt.Println(cars[3])

	// changing an element
	cars[0] = "Toyota"
	fmt.Println(cars)

	arr5 := [5]int{}              // not initialized
	arr6 := [5]int{1, 2}          // partially initialized
	arr7 := [5]int{1, 2, 3, 4, 5} // fully initialized

	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)

	arr8 := [5]int{1: 10, 2: 40} // initializes only the second and third elements of the array

	fmt.Println(arr8)
	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
}
