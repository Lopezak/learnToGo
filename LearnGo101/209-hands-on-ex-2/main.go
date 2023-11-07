package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() string {
	return fmt.Sprintf("Hello! My name is %v %v and I am %v years old!", p.first, p.last, p.age)
}

type human interface {
	speak() string
}

func saySomething(h human) string {
	return h.speak()
}

func main() {
	adam := person{
		first: "Adam",
		last:  "Lopez",
		age:   28,
	}

	kaitlyn := person{
		first: "Kaitlyn",
		last:  "Lopez",
		age:   23,
	}

	// fmt.Println(adam.speak())
	fmt.Println(saySomething(&adam))
	fmt.Println(saySomething(&kaitlyn))

}
