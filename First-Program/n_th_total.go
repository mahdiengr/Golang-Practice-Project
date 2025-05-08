package main

import "fmt"

func n_total(n int) {
	if n <= 100 {
		n = n * (n + 1) / 2
		fmt.Println(n)
	}

}

func main() {
	var n int
	fmt.Scan(&n)
	n_total(n)

}
