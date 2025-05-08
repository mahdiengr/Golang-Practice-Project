package main

import "fmt"

func leap_year_list() {

	for year := 1990; year <= 2024; year++ {
		if (year%4 == 0) && (year%100 != 0 || year%400 == 0) {
			fmt.Println(year)
		}
	}
}

func main() {
	// var year int
	// fmt.Println("Enter a year:")
	// fmt.Scanln(&year)
	// Check if the year is a leap year or not
	leap_year_list()
}
