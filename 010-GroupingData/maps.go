package main

import "fmt"

func main()  {
	fmt.Println("maps.go go ;)")

	m := map[string]int{
		"James": 32,
		"Miss Moneypenny": 27,

	}
	fmt.Println(m)

	fmt.Println(m["James"])

	// If there is no entry in the map it returns zero value

	// ask if value exists comma ok idiom
	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("This is the if Print Miss Moneypennys age:", v)
	}
}