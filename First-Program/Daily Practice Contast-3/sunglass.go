package main

import "fmt"

func main() {
	var x, t int
	fmt.Scan(&x, &t)
	if x >= 1 && x <= 1000000000 && t >= 1 && t <= 1000000000 {
		if t < x {
			fmt.Println(x - t)
		} else {
			fmt.Println(0)
		}
	}
}
