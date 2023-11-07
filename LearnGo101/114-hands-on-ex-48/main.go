package main

import "fmt"

func main() {
	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "I'm 008."}

	c := [][]string{a, b}

	for _, v := range c {
		for _, x := range v {
			fmt.Println(x)
		}
	}
}
