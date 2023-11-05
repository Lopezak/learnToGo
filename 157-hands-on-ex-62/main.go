package main

import (
	"fmt"
	"math"
)

func main() {
	s := Square{
		length: 5,
		width:  8,
	}

	c := Circle{
		radius: 4,
	}

	info(s)
	info(c)
}

type Square struct {
	length float64
	width  float64
}

func (s Square) Area() float64 {
	return s.length * s.width
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type Shape interface {
	Area() float64
}

func info(s Shape) {
	fmt.Printf("The area of the %T is %v\n", s, s.Area())
}
