package main

import "fmt"

var a = 10

func main() {
	fmt.Println("I am main a:", a)

}

func init() {
	fmt.Println("I am Init", a)
	a := 20
	fmt.Println("I am Init", a)
}
