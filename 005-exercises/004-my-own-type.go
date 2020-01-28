package main

import "fmt"

type zerodev int

var x zerodev

func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Println(x)
}
