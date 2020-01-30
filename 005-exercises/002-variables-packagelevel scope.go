package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func two() {

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
