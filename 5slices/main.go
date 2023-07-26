package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	myslice3 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	// with omitted capacity
	myslice4 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))

	prices := []int{10, 20, 30}
	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// append elements to slice
	myslice5 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice5 = %v\n", myslice5)
	fmt.Printf("length = %d\n", len(myslice5))
	fmt.Printf("capacity = %d\n", cap(myslice5))

	myslice5 = append(myslice5, 20, 21)
	fmt.Printf("myslice5 = %v\n", myslice5)
	fmt.Printf("length = %d\n", len(myslice5))
	fmt.Printf("capacity = %d\n", cap(myslice5))

	// The '...' after slice2 is necessary when appending the elements of one slice to another.
	myslice6 := append(myslice1, prices...)
	fmt.Printf("myslice6 = %v\n", myslice6)

	// reslicing
	myslice7 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice7 = %v\n", myslice7)
	fmt.Printf("length = %d\n", len(myslice7))
	fmt.Printf("capacity = %d\n", cap(myslice7))
	myslice7 = myslice7[:3]
	fmt.Printf("myslice7 = %v\n", myslice7)
	fmt.Printf("length = %d\n", len(myslice7))
	fmt.Printf("capacity = %d\n", cap(myslice7))

	// copy slice elements to reduce memory usage
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}
