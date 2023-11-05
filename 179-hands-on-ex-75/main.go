package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop"
	q := "Persistently, patiently"
	r := "The meaning of life is..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("The value is %v of type %T with data %v\n", a, a, *a)
	fmt.Printf("The value is %v of type %T with data %v\n", b, b, *b)
	fmt.Printf("The value is %v of type %T with data %v\n", c, c, *c)
	fmt.Printf("The value is %v of type %T with data %v\n", d, d, *d)
}
