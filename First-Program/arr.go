package main

import "fmt"

func main() {
	x := []int{3, 4, 10, 9, 11, 15}
	// fmt.Println(len(x))
	// fmt.Println(cap(x))
	x = append(x, 10)
	fmt.Println(x)
	y := x[3:6]
	y[2] = 20
	z := x

	fmt.Println("y:", y)
	fmt.Println("x:", x)
	fmt.Println("z:", z)
	z[4] = 30
	z = append(z, 50)
	fmt.Println("x:", x)
	fmt.Println("z:", z)
	z[2] = 22
	fmt.Println("x:", x)

	fmt.Println("z:", z)
	fmt.Println("y:", y)
	// fmt.Println(len(z))
	// fmt.Println(cap(z))
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

}
