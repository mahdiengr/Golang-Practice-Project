package main

import "fmt"

var add = func(x int, b int) {
	c := x + b
	fmt.Println("anynomous:", c)

}

func main() {

	fmt.Println("Main function:")
	add(3, 5)

}

// func init() {
// 	x = 10
// 	fmt.Println("init function:", x)
// }
