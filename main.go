package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	//variable declaration
	fmt.Println(nameOne)
	//strings

	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yash"
	fmt.Println(nameFour)

	//ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits $ memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 25

	var scoreOne float32 = -1.5
	var scoreTwo float64 = 2767322332.2
	scroreThree := 1.4
}
