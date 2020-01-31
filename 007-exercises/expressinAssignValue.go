package main

import "fmt"

func main()  {
	// all true
	a := 1 == 1
	b := 1 <= 1
	c := 1 >= 1
	d := 1 != 2
	e := 1 < 2
	f := 1 > 0

	fmt.Println(a, b, c, d, e, f)
}