package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])

	//looping
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as key type

	phonebook := map[int]string{
		12121: "mario",
		35434: "logit",
		32323: "yash",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[12121])

	phonebook[12121] = "bowser"
	fmt.Println(phonebook)

}
