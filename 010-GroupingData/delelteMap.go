package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	// delete(m, "James")
	// fmt.Println(m)

	// comma ok idiom
	if v, ok := m["James"]; ok {
		fmt.Println("value:", v)
		delete(m, "James")
	}
	fmt.Println("Printing end", m)

}
