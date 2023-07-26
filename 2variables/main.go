package main

import (
	"fmt"
)

func main() {
	var student1 string = "Tomato" //type is string
	var student2 = "Zomato"        //type is inferred
	no := 2                        //type is inferred
	// inferred: means that the compiler decides the type of the variable, based on the value

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(no)

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var student3 string
	student3 = "Potato"
	fmt.Println(student3)

	// multivariable declaration
	var x, y, z int = 1, 2, 3
	fmt.Println(x, y, z)

	// for better readability
	var (
		aa int
		bb int    = 1
		cc string = "hello"
	)

	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	
	// declaring a typed constant
	const PI float64 = 3.14159265358979323846264338327950288419716939937510582097494459
	
	// declaring an untyped constant
	const PI2 = 3.14159265358979323846264338327950288419716939937510582097494459

	fmt.Println(PI)
	fmt.Println(PI2)
}
