package main

import "fmt"

func main() {
	var x, y, r int
	fmt.Scan(&x, &y)

	if (x >= 1 && x <= 9) && (y >= 1 && y <= 9) {

		r = x + y
		if r <= 9 {
			fmt.Println(r)

		} else {
			fmt.Println("error")
		}

	}

}
