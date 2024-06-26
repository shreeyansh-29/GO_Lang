package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good Morning, %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Good Bye, %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// sayGreeting("Mario")
	// sayGreeting("Yash")
	// sayBye("Mario")

	// cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	// cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("Circle 1 is %0.3f \n", a1)
}
