package main

import "fmt"

func main() {
	// Arrays are not recommended in GO USE Slices!
	var x [5]int
	fmt.Println(x)

	x[3] = 42
	fmt.Println(x)

	fmt.Println(len(x))

}
