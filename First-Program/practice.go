package main

import "fmt"

func welcomeMessage() {
	fmt.Println("Welcome to our Site")
}

func getName() string {
	var userName string
	fmt.Println("Enter Ur Name: ")
	fmt.Scanln(&userName)
	return userName
}

func getTwoNumber() (int, int) {
	var x int
	var y int
	fmt.Println("Enter 1st number: ")
	fmt.Scanln(&x)
	fmt.Println("Enter 2nd number: ")
	fmt.Scanln(&y)
	return x, y
}

func add(x int, y int) int {
	return x + y
}

func display(name string, x int) {
	fmt.Println("Hi", name, "Summation Result :", x)
}

func goodbyMessage() {
	fmt.Println("----Good By-----")
}
func main() {

	welcomeMessage()
	name := getName()
	num1, num2 := getTwoNumber()
	sum := add(num1, num2)
	display(name, sum)
	goodbyMessage()

}
