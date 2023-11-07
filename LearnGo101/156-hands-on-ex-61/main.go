package main

import "fmt"

func main() {
	p := Person{
		first: "Adam",
		age:   28,
	}

	fmt.Println(p.speak())
}

type Person struct {
	first string
	age   int
}

func (p Person) speak() string {
	return fmt.Sprintf("My name is %v and I am %v years old", p.first, p.age)
}
