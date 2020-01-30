package main

import "fmt"

type zerodev int

var x zerodev

func four() {
	fmt.Println(x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Println(x)
}
