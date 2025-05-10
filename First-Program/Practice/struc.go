// package main

// import "fmt"

// type User struct {
// 	Name string
// 	Age  int
// }

// func PrintDetails(usr User) {
// 	fmt.Println(usr.Name)
// 	fmt.Println(usr.Age)
// }
// func main() {

// 	user1 := User{
// 		Name: "mahdi",
// 		Age:  31,
// 	}

// 	user2 := User{
// 		Name: "Reaz",
// 		Age:  40,
// 	}

// 	PrintDetails(user1)
// 	PrintDetails(user2)

// }
package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Salary float64
}

func print(numbers [3]int) {
	fmt.Println(numbers)
}

func print2(numbers *[3]int) {
	fmt.Println(numbers)
}

func main() {
	x := 20
	p := &x
	*p = 30

	fmt.Println(x) // 30
	fmt.Println("Address:", p)
	fmt.Println("Value:", *p)

	arr := [3]int{1, 2, 3}
	print(arr)   // pass by value
	print2(&arr) // pass by reference

	user1 := User{
		Name:   "Ruhin",
		Age:    21,
		Salary: 0,
	}
	p2 := &user1
	fmt.Println(p2.Age)
}
