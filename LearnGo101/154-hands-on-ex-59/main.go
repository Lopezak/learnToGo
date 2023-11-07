package main

import "fmt"

func main() {
	defer fmt.Println("This was deferred")
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(3))
	fmt.Println(foo(xi...))

	//fmt.Println(bar(3))
	fmt.Println(bar(xi))
}

func foo(n ...int) int {
	var sum int
	for _, v := range n {
		sum += v
	}
	return sum
}

func bar(n []int) int {
	var sum int
	for _, v := range n {
		sum += v
	}
	return sum
}
