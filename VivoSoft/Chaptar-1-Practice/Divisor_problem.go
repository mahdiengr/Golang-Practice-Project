package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&n)
	sum := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("The sum of all integers from 1 to", n, "that are divisible by either 3 or 5 is", sum)
}
