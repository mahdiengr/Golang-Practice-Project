package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n >= 1 && n <= 1000000000 {
		if n%2 == 0 {
			fmt.Println("Mahmoud")
		} else {
			fmt.Println("Ehab")
		}
	}
}
