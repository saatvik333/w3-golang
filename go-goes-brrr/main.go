package main

import (
	"fmt"
	"learning/modules"
	submo "testing2"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello, World!")
	Gun()
	modules.Test()
	modules.Powder()
	fmt.Println(uuid.NewString())
	submo.Hell()
}
