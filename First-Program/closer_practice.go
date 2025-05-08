package main

import "fmt"

// const a = 10

// var p = 100

//Closure
func outer() func() {
	// age := 30
	// fmt.Println("Age =", age)

	// show := func() {
	// 	money = money + a + p
	// 	fmt.Println(money)
	// }

	// return show
	count := 0

	c := func() {
		count = count + 1
		fmt.Println(count)

	}
	return c

}

func call() {
	// incr1 := outer()
	// incr1()
	// inc2 := outer()
	// inc2()
	// incr1 := outer(100)
	// incr1() // money = 100 + 10 + 100 = 210
	// incr1() // money = 210 + 10 + 100 = 320

	// incr2 := outer(100)
	// incr2()
	// incr2()
	// incr1()
}

func main() {
	// call()
	ca := outer()
	ca()
	test := func() {
		fmt.Println("test")
	}

	test()
}

func init() {
	fmt.Println("=== Bank ===")
}
