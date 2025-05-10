package main

import "fmt"

func main() {
	var t, x int
	fmt.Scan(&t)
	if t <= 10 {
		for i := 1; i <= t; i++ {
			fmt.Scan(&x)
			if x <= 10 {
				x = x * 2 * 5
				fmt.Println(x)
			}

		}

	}

}
