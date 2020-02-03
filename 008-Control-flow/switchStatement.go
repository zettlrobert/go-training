package main

import "fmt"

func main() {

	// switch can have a value case case == value --> executes
	switch {
	case false:
		fmt.Println("this should not print")

	case (2 == 4):
		fmt.Println("this should not print")

	case (3 == 3):
		fmt.Println("prints")
		// fallthrough // everything after gets printed DONT USE

	case (4 == 4):
		// does not only if FALLTHOURGH
		fmt.Println("also true does it print")

	default:
		fmt.Println("This is default")
	}
}
