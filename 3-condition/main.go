package main

import "fmt"

func main() {
	fmt.Print("Input a fruit: ")
	var fruit string
	fmt.Scanln(&fruit)

	if fruit == "" {
		fmt.Println("meh!")
		return
	}

	switch fruit {
	case "apple":
		fmt.Println("APPLE")
	case "banana":
		fmt.Println("BANANA")
	case "lemon":
		fmt.Println("LEMON")
	default:
		fmt.Println("OMG!")
	}
}