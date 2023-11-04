package main

import "fmt"

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {
	adam := person{
		first:    "Adam",
		last:     "Lopez",
		iceCream: []string{"Mint"},
	}

	kailtyn := person{
		first:    "Kaitlyn",
		last:     "obdam",
		iceCream: []string{"Birthday Cake", "Vanilla"},
	}

	xm := make(map[string]person)

	xm[adam.last] = adam
	xm[kailtyn.last] = kailtyn

	for _, v := range xm {
		fmt.Printf("%v %v's favorite ice creams are ", v.first, v.last)
		for _, ic := range v.iceCream {
			fmt.Printf("%v ", ic)
		}
		fmt.Println()
	}

	// fmt.Println(adam.first, adam.last)
	// for _, ic := range adam.iceCream {
	// 	fmt.Println(ic)
	// }

	// fmt.Println(kailtyn.first, kailtyn.last)
	// for _, ic := range kailtyn.iceCream {
	// 	fmt.Println(ic)
	// }
}
