package main

import "fmt"

var input string
var colorCode string

func main() {
	// 	input := 4
	// 	switch input {
	// 	case 1:
	// 		fmt.Println("one")
	// 	case 2:
	// 		fmt.Println("two")
	// 	case 3:
	// 		fmt.Println("three")
	// 	default:
	// 		fmt.Println("invalid value")
	// 	}
	fmt.Scan(&input)
	switch input {
	case "blue":
		fmt.Println("#000FF")
	case "green":
		fmt.Println("#DDWWWP")
	case "pink":
		fmt.Println("SAAWEQ")
	case "yellow":
		fmt.Println("PPOPIIP")

	default:
		fmt.Println("invalid code")
	}
}
