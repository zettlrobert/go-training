package main

import "fmt"

func main() {
	fmt.Println("Hey")

	// slice of string

	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	// fmt.Println(xs1)
	// fmt.Println(xs2)

	// slice of slice of string
	xxs := [][]string{xs1, xs2}

	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}
