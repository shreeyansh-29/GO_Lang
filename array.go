package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{1, 2, 3}

	var ages1 = [3]int{20, 25, 30}

	names := [4]string{"yash", "mario", "peach", "bowser"}

	fmt.Println(ages, len(ages))

	//slices (use arrays under the hood)
	var scores = []int{100, 50, 60}

	scores[2] = 35

	scores = append(scores, 85)

	//slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne)
}
