package main

import "fmt"

func main() {

	fmt.Println("hello, learning go programming language")
	foo()

	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
}

func foo() {
	fmt.Println("I'm in foo")
}

// controlf low:
// (1) sequence
// (2) loop; iterative
// (3) conditional
