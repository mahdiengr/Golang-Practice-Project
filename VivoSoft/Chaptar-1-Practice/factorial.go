package main

import "fmt"

func recursion(n int) int {
	if n == 0 {
		return 1
	}

	return recursion(n-1) * n
}

func main() {
	var n int
	n = 5
	fmt.Println(recursion(n))
}
