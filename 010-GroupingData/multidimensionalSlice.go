package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Choco", "martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	// Multidimensional Slice
	xp := [][]string{jb, mp}
	fmt.Println(xp)

}
