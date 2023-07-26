package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	// using make
	var c = make(map[string]string) // The map is empty now
	c["brand"] = "Ford"
	c["model"] = "Mustang"
	c["year"] = "1964"

	// a is no longer empty
	d := make(map[string]int)
	d["Oslo"] = 1
	d["Bergen"] = 2
	d["Trondheim"] = 3

	fmt.Printf("c\t%v\n", c)
	fmt.Printf("d\t%v\n", d)
	fmt.Printf(a["brand"])

	fmt.Println(a)
	a["year"] = "1970" // Updating an element
	a["color"] = "red" // Adding an element
	delete(a,"year")   // Deleting an element
	fmt.Println(a)
}
