package main

import "fmt"

func main() {
	var char string

	fmt.Scan(&char)

	switch char {
	case "a":
		fmt.Println("vowel")
	case "e":
		fmt.Println("vowel")
	case "i":
		fmt.Println("vowel")
	case "o":
		fmt.Println("vowel")
	case "u":
		fmt.Println("vowel")
	default:
		fmt.Println("consonant")
	}

}
