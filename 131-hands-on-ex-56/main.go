package main

import "fmt"

func main() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Adam",
		friends: map[string]int{
			"Jonny": 28,
			"Mark":  30,
			"Etahn": 26,
		},
		favDrinks: []string{
			"Lemonade",
			"Coffee",
			"Water",
			"Tea",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.friends)
	fmt.Println(p1.favDrinks)
}
