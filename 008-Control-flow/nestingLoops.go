package main

import "fmt"

func main() {
	fmt.Println("Hello -> fmt without reformatting")

	for i := 0; i < 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)
		}
	}

	fmt.Println("done")
}
