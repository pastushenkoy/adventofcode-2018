package main

import (
	"fmt"
	"github.com/pastushenkoy/adventOfCode/utils"
)

func main() {
	data := utils.ReadFileOfStrings("input.txt")

	field := LoadField(data)

	lastCartPosition :=  field.iterateToTheEnd()

	fmt.Println(field.Collisions)

	fmt.Printf("The first collision point is %v\n", field.Collisions[0])
	fmt.Printf("Position of the last cart is %v\n", lastCartPosition)
}


