package main

import "fmt"

func main() {
	var n, result, x, y int
	fmt.Scan(&n)
	if n >= 1 && n <= 100 {
		result = n / 15
		x = n * 800
		y = result * 200
		y = x - y
		fmt.Println(y)

	}
}
