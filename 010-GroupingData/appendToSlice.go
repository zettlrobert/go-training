package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 42}
	fmt.Println(x)

	// append --> đrom built in package
	x = append(x, 12, 13, 14, 15)
	fmt.Println(x)

	y := []int{234, 564, 678, 987}

	// y... puts all values in slice right there
	// ...y variadic take unlimited number of values from this type
	x = append(x, y...)
	fmt.Println(x)


}
