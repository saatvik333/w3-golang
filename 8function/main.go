package main

import (
	"fmt"
)

func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func myFunction(x int, y int) (result int) {
	result = x + y
	return result
}

func myFunction2(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	familyName("Liam", 3)
	familyName("Jenny", 14)
	familyName("Anja", 30)

	// store and print the return value
	total := myFunction(1, 2)
	fmt.Println(total)

	a, b := myFunction2(5, "Hello")
	fmt.Println(a, b)

	_, c := myFunction2(5, "Hello")
	fmt.Println(c)

	testcount(1)
	fmt.Println(factorial_recursion(4))
}
