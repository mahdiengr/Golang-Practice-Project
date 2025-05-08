package mathlib

import "fmt"

func Subtraction(x, y int) {
	if x > y {
		fmt.Println("Subtraction:", x-y)
	} else {
		fmt.Println("Subtraction:", y-x)
	}

}
