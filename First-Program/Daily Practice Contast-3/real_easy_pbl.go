package main

import "fmt"

func main() {
	var t, n int
	fmt.Scan(&t)
	if t >= 1 && t <= 99 {
		for i := 1; i <= t; i++ {
			fmt.Scan(&n)
			if n >= 2 && n <= 100 {
				n = n - 1
				fmt.Println(n)
			}

		}

	}

}
