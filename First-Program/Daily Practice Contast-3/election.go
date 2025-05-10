package main

import "fmt"

func main() {
	var t, n, a, r int
	fmt.Scan(&n, &t, &a)
	r = t + a
	if n >= 1 && n <= 99 {
		if n%2 != 0 && t >= 0 && a >= 0 && t <= n && a <= n && r <= n {
			if t > a {
				if (n-r)+a < t {
					fmt.Println("Yes")
				} else {
					fmt.Println("No")
				}
			} else {
				if (n-r)+t < a {
					fmt.Println("Yes")
				} else {
					fmt.Println("No")
				}

			}

		} else {
			fmt.Println("No")
		}

	}

}
