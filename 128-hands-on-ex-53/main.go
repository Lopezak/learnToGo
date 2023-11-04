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
		last:     "Lopez",
		iceCream: []string{"Birthday Cake", "Vanilla"},
	}

	fmt.Println(adam.first, adam.last)
	for _, ic := range adam.iceCream {
		fmt.Println(ic)
	}

	fmt.Println(kailtyn.first, kailtyn.last)
	for _, ic := range kailtyn.iceCream {
		fmt.Println(ic)
	}
}
