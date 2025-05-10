package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if (a >= -100 && a <= 100) && (b >= -100 && b <= 100) && (c >= -100 && c <= 100) {
		if a <= c && b >= c {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}
