package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 8, 5, 42}
	fmt.Println(x)

	x = append(x, 77, 123, 44, 1203)
	fmt.Println(x)

	y := []int{234, 456, 232, 122}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)

}
