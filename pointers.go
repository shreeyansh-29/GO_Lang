package main

import (
	"fmt"
)

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	name := "tifa"

	// updateName(name)

	m := &name

	fmt.Println("memory address of name is: ", m)
	fmt.Println("value at memory address: ", *m)

	updateName(m)

	fmt.Println(name)
}
