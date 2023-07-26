package main

import (
	"fmt"
)

func main() {
	var i, j string = "Hello", "World"

	fmt.Print(i)
	fmt.Print(j)

	// The Println() function is similar to Print() with the difference that a
	// whitespace is added between the arguments, and a newline is added at the end:
	fmt.Println()
	fmt.Println(i)
	fmt.Println(j)

	// The Printf() function first formats its argument based on the given formatting verb and then prints them.
	var a string = "Hello"
	var b int = 15
	// For more knowledge about formatting verbs https://pkg.go.dev/fmt
	// %v is used to print the value of the arguments
	// %T is used to print the type of the arguments
	fmt.Printf("a has value: %v and type: %T \n", a, a)
	fmt.Printf("b has value: %v and type: %T \n", b, b)

}
