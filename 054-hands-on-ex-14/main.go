package main

import (
	"fmt"
)

var x = 40

const y = 41

func main() {
	fmt.Printf("the value of z is %v and the type of z is %T\n", x, x)
	fmt.Printf("the value of z is %v and the type of z is %T\n", y, y)
	z := 42
	fmt.Printf("the value of z is %v and the type of z is %T\n", z, z)
}

func insideFunction() {
	fmt.Println("hello from the inside")
}
