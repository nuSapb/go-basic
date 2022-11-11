package main

import "fmt"

func main() {

	// 1D Array
	var a [5]int
	a[0] = 10
	a[2] = 30
	a[3] = 40
	// fmt.Println(a)
	// fmt.Println(len(a))
	for _, v := range a {
		fmt.Println(v)
	}

	b := [3]int{1, 2, 3}
	for _, v := range b {
		fmt.Println(v)
		
	}

	// 2D Array
	fmt.Println("2D Array")
	var c [2][3]int
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			c[i][j] = j
		}
		fmt.Println(c)
	}

	// Slide
	fmt.Println("--- Slide ---")
	d := make([]int, 5)
	d[0] = 10
	d[2] = 20
	d[3] = 40
	d = append(d, 90)
	fmt.Println(len(d))
	fmt.Println(d)

	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[:2]
	f[0] = 99
	fmt.Println(e)
	fmt.Println(f)

}