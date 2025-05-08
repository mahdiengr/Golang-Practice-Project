package main

import "fmt"

func main() {
	var day int
	fmt.Println("Enter a day:")
	fmt.Scanln(&day)
	dayName := ""
	switch day {
	case 1:
		dayName = "Sunday"
	case 2:
		dayName = "Monday"
	case 3:
		dayName = "Tuesday"
	case 4:
		dayName = "Wednesday"
	case 5:
		dayName = "Thursday"
	case 6:
		dayName = "Friday"
	case 7:
		dayName = "Saturday"

	}
	fmt.Println("Name of the day is ", dayName)
}
