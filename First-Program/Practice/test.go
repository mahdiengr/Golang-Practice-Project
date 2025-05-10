// package main

// import "fmt"

// var x = 100
// var add = func(x int, b int) {
// 	c := x + b
// 	fmt.Println("anynomous:", c)

// }

// func main() {

// 	fmt.Println("Main function:")
// 	add(3, 5)

// }

// func init() {
// 	x := 10
// 	fmt.Println("init function:", x)
// }

// func init() {
// 	x := 10
// 	fmt.Println("init function:", x)
// }

package main

import "fmt"

func print_disp(arr *[3]string) {
	for _, val := range arr {
		fmt.Println(val)
	}

}
func main() {
	fruits := [3]string{"apple", "orange", "banana"}
	// for _, val := range fruits {
	// 	fmt.Println(val)
	// }

	print_disp(&fruits)
}
