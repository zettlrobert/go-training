package main

import "fmt"

func main()  {
	x := []int{1 ,2, 3, 4, 5, 42}
	fmt.Println(x)

	// slicing a slice
	fmt.Println(x[1:])
	// slice beginning position upto but not including
	fmt.Println(x[1:3])


	for i := 0; i < len(x); i++ {
		fmt.Println("Printing Slice on Position:", i, x[i])
	}
}