package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
		`fleming_ian`:     []string{`steaks`, `cigars`, `espionage`},
	}

	// delete
	delete(m, `no_dr`)

	for k, val := range m {
		fmt.Println("this is the record for", k)
		for i, val2 := range val {
			fmt.Println("\t", i, val2)
		}
	}

}
