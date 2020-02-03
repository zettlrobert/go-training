package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 8, 5, 42}
	fmt.Println(x)

	// if you know size of slize you can use make
	// slice sits ontop of array

	// type, lengbth, capacity
	y := make([]int, 10, 100)

	y = append(y, 3434)

	fmt.Println(len(y))
	fmt.Println(cap(y))

	// if capacity raches size, a new arra with double the capacity is created
}
