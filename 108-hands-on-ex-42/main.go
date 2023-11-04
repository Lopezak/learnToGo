package main

import "fmt"

func main() {
	a := [5]int{0, 1, 2, 3, 4}

	for i, v := range a {
		fmt.Printf("Position %v is %v with type %T\n", i, v, v)
	}
}
