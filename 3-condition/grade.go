package main

import (
	"fmt"
)

func main() {
	fmt.Print("Input score: ")
	var score int
	fmt.Scanln(&score)

	// if else example

	// if score < 50 {
	// 	fmt.Println("F")
	// } else if score < 60 {
	// 	fmt.Println("D")
	// } else if score < 70 {
	// 	fmt.Println("C")
	// } else if score < 80 {
	// 	fmt.Println("B")
	// } else {
	// 	fmt.Println("S")
	// }

	// switch case example 
	switch {
	case score < 50:
		fmt.Println("F")
	case score < 60:
		fmt.Println("D")
	case score < 70:
		fmt.Println("C")
	case score < 80:
		fmt.Println("B")
	default:
		fmt.Println("S")
	}
} 