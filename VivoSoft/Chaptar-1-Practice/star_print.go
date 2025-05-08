package main

import (
	"fmt"
)

func main() {
	// var n int
	for i := 1; i <= 6; i++ {
		fmt.Print(i, ":")
		for j := 1; j <= i; j++ {
			fmt.Print("**")
		}
		fmt.Println()
	}
}
