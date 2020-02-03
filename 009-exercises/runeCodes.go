package main

import "fmt"

func main() {
	// Print every rune code point of the uppercase alphabet three times.
	// ACCI 65 to 90

	for i := 65; i <= 90; i++ {
		fmt.Printf("%#U\n", i)
	}

}
