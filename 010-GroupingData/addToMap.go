package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	// Adding Tod with value 33 to Map
	m["Todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 6, 7, 8, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}

}
