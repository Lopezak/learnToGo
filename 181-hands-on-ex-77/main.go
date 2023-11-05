package main

import "fmt"

type Person struct {
	first string
}

func valSem(p Person, s string) Person {
	p.first = s
	return p
}

func poiSem(p *Person, s string) {
	p.first = s
}

func main() {
	p := Person{
		first: "Adam",
	}

	fmt.Println(p.first)
	p = valSem(p, "Kaitlyn")
	fmt.Println(p.first)
	poiSem(&p, "Ellie")
	fmt.Println(p.first)
}
