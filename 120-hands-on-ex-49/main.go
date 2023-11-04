package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:       {`shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`: {`intelligence`, `literature`, `computer science`},
		`no_dr`:            {`cats`, `ice cream`, `sunsets`},
	}

	// for k, v := range x {
	// 	fmt.Println(k)
	// 	for i, v2 := range v {
	// 		fmt.Printf("Index is %v and value is %v\n", i, v2)
	// 	}
	// }

	x[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	// for k, v := range x {
	// 	fmt.Println(k)
	// 	for i, v2 := range v {
	// 		fmt.Printf("Index is %v and value is %v\n", i, v2)
	// 	}
	// }

	delete(x, `no_dr`)

	for k, v := range x {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Printf("Index is %v and value is %v\n", i, v2)
		}
	}

}
