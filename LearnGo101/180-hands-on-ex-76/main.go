package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Printf("%v: I'm walkin over here!\n", d.first)
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	youngRun(d1)

	d2 := dog{"Padget"}
	youngRun(d2)
}
