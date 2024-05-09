package main

import "fmt"

func main() {
	// x := 0

	// for x < 5 {
	// 	fmt.Println("value of x:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i:", i)
	// }

	names := [4]string{"yash", "mario", "peach", "bowser"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Println(value, index)
	// }

	for _, value := range names {
		fmt.Println(value)
	}
}
