package main

import "fmt"

func getNumber() int {
	var num int
	fmt.Print("Enter ur Number:")
	fmt.Scanln(&num)
	return num
}

// func loop_print(n int) {

// 	for i := 1; i <= n; i++ {
// 		fmt.Println(i)
// 	}

// }

func arrar_insert(n int, arr *[]int) {

	for i := 0; i < n; i++ {
		fmt.Scan(&(*arr)[i])
	}

}

func PrintArray(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		fmt.Print((*arr)[i], " ")
	}
}

// func PrintArray(arr *[]int) {
// 	for _, v := range *arr{
// 		fmt.Println(v)

// 	}
// }

func main() {
	x := getNumber()
	// // loop_print(x)
	arr := make([]int, x)
	arrar_insert(x, &arr)

	// PrintArray(&arr)

	// & (bitwise AND)
	sum := 0
	for _, v := range arr {
		if v&1 != 0 {
			sum++
		}

	}

	fmt.Println(sum)
	// for i := 1; i <= 10; i++ {
	// 	x := y << i
	// 	fmt.Println(y, "*", i, "=", x)
	// }
	// fmt.Println("after right shift")
	// for i := 1; i <= 10; i++ {
	// 	x := y >> i
	// 	fmt.Println(y, "*", i, "=", x)
	// }

}
